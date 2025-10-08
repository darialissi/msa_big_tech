package usecases

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/mocks"
)

type AcceptFriendRequestTestSuite struct {
	suite.Suite

	RepoFriend *mocks.FriendRepository
	RepoFriendReq *mocks.FriendRequestRepository

	Usecase *SocialUsecase
}

// SetupTest выполняется перед каждым тестом
func (s *AcceptFriendRequestTestSuite) SetupTest() {
	s.RepoFriendReq = mocks.NewFriendRequestRepository(s.T())
	s.RepoFriend = mocks.NewFriendRepository(s.T())

    // Инициализируем Usecase перед установкой зависимостей
    s.Usecase = &SocialUsecase{
        Deps: Deps{
            RepoFriendReq: s.RepoFriendReq,
            RepoFriend:    s.RepoFriend,
        },
    }
	// конструктор
}

// TearDownTest выполняется после каждого теста
func (s *AcceptFriendRequestTestSuite) TearDownTest() {
	s.RepoFriendReq.AssertExpectations(s.T())
	s.RepoFriend.AssertExpectations(s.T())
	// деструктор
}

func TestAcceptFriendRequestTestSuite(t *testing.T) {
	suite.Run(t, new(AcceptFriendRequestTestSuite))
}

func (s *AcceptFriendRequestTestSuite) Test_AcceptFriendRequest_Positive() {
	// ARRANGE
	var (
		mockCtx = &mocks.MockContext{}

		REQ_ID = dto.FriendRequestID(uuid.New().String())
		FROM_USER_ID = dto.UserID(uuid.New().String())
		TO_USER_ID = dto.UserID(uuid.New().String())

		expected = &models.FriendRequest{
			ID: models.FriendRequestID(REQ_ID),
			Status: models.FriendRequestStatusAccepted,
			FromUserID: models.UserID(FROM_USER_ID),
			ToUserID: models.UserID(TO_USER_ID),
		}
	)

	s.RepoFriendReq.EXPECT().
		FetchById(mockCtx, REQ_ID).
		Return(&models.FriendRequest{
			Status: models.FriendRequestStatusPending,
			FromUserID: models.UserID(FROM_USER_ID),
			ToUserID: models.UserID(TO_USER_ID),
			}, 
			nil).
		Once()

	s.RepoFriendReq.EXPECT().
		UpdateStatus(mockCtx, &dto.UpdateFriendRequest{
			ReqID: REQ_ID,
			Status: dto.FriendRequestStatus(models.FriendRequestStatusAccepted),
	}).
		Return(&models.FriendRequest{
			Status: models.FriendRequestStatusAccepted,
			FromUserID: models.UserID(FROM_USER_ID),
			ToUserID: models.UserID(TO_USER_ID),
			}, 
			nil).
		Once()

	s.RepoFriend.EXPECT().
		Save(mockCtx, &dto.SaveFriend{
			UserID: FROM_USER_ID,
			FriendID: TO_USER_ID,
	}).
		Return(&models.UserFriend{
			UserID: models.UserID(FROM_USER_ID),
			FriendID: models.UserID(TO_USER_ID),
			}, 
			nil).
		Once()

	// ACT

	got, err := s.Usecase.AcceptFriendRequest(mockCtx, REQ_ID)

	// ASSERT
	s.NoError(err)
	s.NotNil(got)

	s.Equal(expected.FromUserID, got.FromUserID)
	s.Equal(expected.ToUserID, got.ToUserID)
	s.Equal(expected.Status, got.Status)
}