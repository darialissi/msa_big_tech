package usecases

import (
	"testing"

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
		reqId = "11111111-0000-0000-0000-0000000000000"

		REQ_ID = dto.FriendRequestID(reqId)
		FROM_USER_ID = dto.UserID("00000000-0000-0000-0000-0000000000000")
		TO_USER_ID = dto.UserID("00000000-0000-0000-0000-0000000000001")

		expected = &models.FriendRequest{
			ID: "10000000-0000-0000-0000-0000000000000",
			Status: models.FriendRequestStatusAccepted,
			FromUserID: "00000000-0000-0000-0000-0000000000000",
			ToUserID: "00000000-0000-0000-0000-0000000000001",
		}
	)

	s.RepoFriendReq.EXPECT().
		FetchById(REQ_ID).
		Return(&models.FriendRequest{
			ID: models.FriendRequestID(REQ_ID),
			Status: models.FriendRequestStatusPending,
			FromUserID: models.UserID(FROM_USER_ID),
			ToUserID: models.UserID(TO_USER_ID),
			}, 
			nil).
		Once()

	s.RepoFriendReq.EXPECT().
		UpdateStatus(&dto.ChangeStatus{
			ReqID: REQ_ID,
			StatusID: dto.FriendRequestStatus(models.FriendRequestStatusAccepted),
	}).
		Return(&models.FriendRequest{
			ID: models.FriendRequestID(REQ_ID),
			Status: models.FriendRequestStatusAccepted,
			FromUserID: models.UserID(FROM_USER_ID),
			ToUserID: models.UserID(TO_USER_ID),
			}, 
			nil).
		Once()

	// ACT
	got, err := s.Usecase.AcceptFriendRequest(dto.FriendRequestID(reqId))

	// ASSERT
	s.NoError(err)
	s.NotNil(got)
	s.NotEmpty(got.ID) // ID должен быть выставлен

	// для сравнения обнулим ID у обоих
	got.ID = ""
	expected.ID = ""
	s.Equal(expected, got)
}