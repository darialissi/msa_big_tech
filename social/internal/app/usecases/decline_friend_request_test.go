package usecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/mocks"
)

func Test_DeclineFriendRequest_whitebox_mockery(t *testing.T) {
	t.Parallel()

	FROM_USER_ID := dto.UserID("00000000-0000-0000-0000-0000000000000")
	TO_USER_ID := dto.UserID("00000000-0000-0000-0000-0000000000001")
	REQ_ID := dto.FriendRequestID("11111111-0000-0000-0000-0000000000000")

	// ARRANGE
	type args struct {
		reqId dto.FriendRequestID
	}
	tests := []struct {
		name      string
		args      args
		mock      func(t *testing.T) Deps
		want      *models.FriendRequest
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test 1. Positive",
			args: args{
				reqId: REQ_ID,
			},
			mock: func(t *testing.T) Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)

				frReqMock.EXPECT().
					FetchById(REQ_ID).
					Return(&models.FriendRequest{
						ID: models.FriendRequestID(REQ_ID),
						Status: models.FriendRequestStatusPending,
						FromUserID: models.UserID(FROM_USER_ID),
						ToUserID: models.UserID(TO_USER_ID),
					}, nil).
					Once()

				frReqMock.EXPECT().
					UpdateStatus(&dto.ChangeStatus{
						ReqID: REQ_ID,
						StatusID: dto.FriendRequestStatus(models.FriendRequestStatusDeclined),
						}).
					Return(&models.FriendRequest{
						ID: models.FriendRequestID(REQ_ID),
						Status: models.FriendRequestStatusDeclined,
						FromUserID: models.UserID(FROM_USER_ID),
						ToUserID: models.UserID(TO_USER_ID),
						}, nil).
					Once()

				frMock := mocks.NewFriendRepository(t)

				return Deps{
					RepoFriendReq: frReqMock,
					RepoFriend: frMock,
				}
			},
			want: &models.FriendRequest{
				Status: models.FriendRequestStatusDeclined,
				FromUserID: models.UserID(FROM_USER_ID),
				ToUserID: models.UserID(TO_USER_ID),
			},
			assertErr: assert.NoError,
		},
		{
			name: "Test 2. Negative: RepoFriendReq.FetchById error",
			args: args{
				reqId: REQ_ID,
			},
			mock: func(t *testing.T) Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)

				frReqMock.EXPECT().
					FetchById(REQ_ID).
					Return(nil, errors.New("RepoFriendReq.FetchById error")).
					Once()

				frMock := mocks.NewFriendRepository(t)

				return Deps{
					RepoFriendReq: frReqMock,
					RepoFriend: frMock,
				}
			},
			want: nil,
			assertErr: assert.Error,
		},
		{
			name: "Test 3. Negative: Already Accepted Request",
			args: args{
				reqId: REQ_ID,
			},
			mock: func(t *testing.T) Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)

				frReqMock.EXPECT().
					FetchById(REQ_ID).
					Return(&models.FriendRequest{
						ID: models.FriendRequestID(REQ_ID),
						Status: models.FriendRequestStatusAccepted,
						FromUserID: models.UserID(FROM_USER_ID),
						ToUserID: models.UserID(TO_USER_ID),
					}, nil).
					Once()

				frMock := mocks.NewFriendRepository(t)

				return Deps{
					RepoFriendReq: frReqMock,
					RepoFriend: frMock,
				}
			},
			want: nil,
			assertErr: assert.Error,
		},
		{
			name: "Test 4. Negative: Already Declined Request",
			args: args{
				reqId: REQ_ID,
			},
			mock: func(t *testing.T) Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)

				frReqMock.EXPECT().
					FetchById(REQ_ID).
					Return(&models.FriendRequest{
						ID: models.FriendRequestID(REQ_ID),
						Status: models.FriendRequestStatusDeclined,
						FromUserID: models.UserID(FROM_USER_ID),
						ToUserID: models.UserID(TO_USER_ID),
					}, nil).
					Once()

				frMock := mocks.NewFriendRepository(t)

				return Deps{
					RepoFriendReq: frReqMock,
					RepoFriend: frMock,
				}
			},
			want: nil,
			assertErr: assert.Error,
		},
		{
			name: "Test 5. Negative: RepoFriendReq.UpdateStatus error",
			args: args{
				reqId: REQ_ID,
			},
			mock: func(t *testing.T) Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)

				frReqMock.EXPECT().
					FetchById(REQ_ID).
					Return(&models.FriendRequest{
						ID: models.FriendRequestID(REQ_ID),
						Status: models.FriendRequestStatusPending,
						FromUserID: models.UserID(FROM_USER_ID),
						ToUserID: models.UserID(TO_USER_ID),
					}, nil).
					Once()

				frReqMock.EXPECT().
					UpdateStatus(&dto.ChangeStatus{
						ReqID: REQ_ID,
						StatusID: dto.FriendRequestStatus(models.FriendRequestStatusDeclined),
						}).
					Return(nil, errors.New("RepoFriendReq.UpdateStatus error")).
					Once()

				frMock := mocks.NewFriendRepository(t)

				return Deps{
					RepoFriendReq: frReqMock,
					RepoFriend: frMock,
				}
			},
			want: nil,
			assertErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// ARRANGE
			deps := tt.mock(t) // создаем мок объекты
			uc := &SocialUsecase{    // инициализируем Usecase
				Deps: deps,
			}

			// ACT
			got, err := uc.DeclineFriendRequest(tt.args.reqId)

			// ASSERT
			tt.assertErr(t, err)

			if got != nil {
				assert.NotEmpty(t, got.ID)
				got.ID = ""
			}

			assert.Equal(t, tt.want, got)  // mark test as failed but continue execution
			// require.Equal(t, tt.want, got) // mark test as failed and exit
		})
	}
}