package usecases

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/mocks"
)

type AcceptFriendRequestTestSuite struct {
	suite.Suite

	RepoFriend    *mocks.FriendRepository
	RepoFriendReq *mocks.FriendRequestRepository
	TxMan         *mocks.TxManager

	Usecase *SocialUsecase
}

// SetupTest выполняется перед каждым тестом
func (s *AcceptFriendRequestTestSuite) SetupTest() {
	s.RepoFriendReq = mocks.NewFriendRequestRepository(s.T())
	s.RepoFriend = mocks.NewFriendRepository(s.T())
	s.TxMan = mocks.NewTxManager(s.T())

	// Инициализируем Usecase перед установкой зависимостей
	s.Usecase = &SocialUsecase{
		Deps: Deps{
			RepoFriendReq: s.RepoFriendReq,
			RepoFriend:    s.RepoFriend,
			TxMan:         s.TxMan,
		},
	}
	// конструктор
}

// TearDownTest выполняется после каждого теста
func (s *AcceptFriendRequestTestSuite) TearDownTest() {
	s.RepoFriendReq.AssertExpectations(s.T())
	s.RepoFriend.AssertExpectations(s.T())
	s.TxMan.AssertExpectations(s.T())
	// деструктор
}

func TestAcceptFriendRequestTestSuite(t *testing.T) {
	suite.Run(t, new(AcceptFriendRequestTestSuite))
}

func (s *AcceptFriendRequestTestSuite) Test_AcceptFriendRequest_Positive() {
	// ARRANGE
	var (
		mockCtx = &mocks.MockContext{}

		REQ_ID       = models.FriendRequestID(uuid.New().String())
		FROM_USER_ID = models.UserID(uuid.New().String())
		TO_USER_ID   = models.UserID(uuid.New().String())

		expected = &models.FriendRequest{
			ID:         REQ_ID,
			Status:     models.FriendRequestStatusAccepted,
			FromUserID: FROM_USER_ID,
			ToUserID:   TO_USER_ID,
		}
	)

	s.RepoFriendReq.EXPECT().
		FetchById(mockCtx, REQ_ID).
		Return(&models.FriendRequest{
			ID:         REQ_ID,
			Status:     models.FriendRequestStatusPending,
			FromUserID: FROM_USER_ID,
			ToUserID:   TO_USER_ID,
		},
			nil).
		Once()

	// TxMan должен быть вызван первым в транзакции
	s.TxMan.EXPECT().
		RunReadCommitted(mockCtx, mock.MatchedBy(func(fn func(txCtx context.Context) error) bool {
			// Внутри транзакции ожидаются вызовы UpdateStatus и Save
			mockTxCtx := &mocks.MockContext{}

			s.RepoFriendReq.EXPECT().
				UpdateStatus(mockTxCtx, &models.FriendRequest{
					ID:     REQ_ID,
					Status: models.FriendRequestStatusAccepted,
				}).
				Return(&models.FriendRequest{
					ID:         REQ_ID,
					Status:     models.FriendRequestStatusAccepted,
					FromUserID: FROM_USER_ID,
					ToUserID:   TO_USER_ID,
				}, nil).
				Once()

			s.RepoFriend.EXPECT().
				Save(mockTxCtx, &models.UserFriend{
					UserID:   FROM_USER_ID,
					FriendID: TO_USER_ID,
				}).
				Return(&models.UserFriend{
					UserID:   FROM_USER_ID,
					FriendID: TO_USER_ID,
				}, nil).
				Once()

			// Выполняем функцию транзакции
			err := fn(mockTxCtx)
			return err == nil
		})).
		Return(nil).
		Once()

	// ACT

	got, err := s.Usecase.AcceptFriendRequest(mockCtx, dto.FriendRequestID(REQ_ID))

	// ASSERT
	s.NoError(err)
	s.NotNil(got)

	s.Equal(expected.ID, got.ID)
	s.Equal(expected.FromUserID, got.FromUserID)
	s.Equal(expected.ToUserID, got.ToUserID)
	s.Equal(expected.Status, got.Status)
}
