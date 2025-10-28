package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	uc "github.com/darialissi/msa_big_tech/social/internal/app/usecases"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/tests/mocks"
)

func Test_AcceptFriendRequest_whitebox_mockery(t *testing.T) {
	t.Parallel()

	mockCtx := &mocks.MockContext{}
	mockTxCtx := &mocks.MockContext{}

	REQ_ID := models.FriendRequestID(uuid.New().String())
	FROM_USER_ID := models.UserID(uuid.New().String())
	TO_USER_ID := models.UserID(uuid.New().String())

	// ARRANGE
	type args struct {
		reqId dto.FriendRequestID
	}
	tests := []struct {
		name      string
		args      args
		mock      func(t *testing.T) uc.Deps
		want      *models.FriendRequest
		assertErr assert.ErrorAssertionFunc
	}{
		{
			name: "Test 1. Positive",
			args: args{
				reqId: dto.FriendRequestID(REQ_ID),
			},
			mock: func(t *testing.T) uc.Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)
				frMock := mocks.NewFriendRepository(t)
				bxMock := mocks.NewOutboxRepository(t)
				txManMock := mocks.NewTxManager(t)

				frReqMock.EXPECT().
					FetchById(mockCtx, REQ_ID).
					Return(&models.FriendRequest{
						ID:         REQ_ID,
						Status:     models.FriendRequestStatusPending,
						FromUserID: FROM_USER_ID,
						ToUserID:   TO_USER_ID,
					}, nil).
					Once()

				txManMock.EXPECT().
					RunReadCommitted(mockCtx, mock.AnythingOfType("func(context.Context) error")).
					Run(func(ctx context.Context, fn func(context.Context) error) {
						// Имитируем выполнение внутри транзакции
						fn(mockTxCtx)
					}).
					Return(nil).
					Once()

				frReqMock.EXPECT().
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

				frMock.EXPECT().
					Save(mockTxCtx, &models.UserFriend{
						UserID:   FROM_USER_ID,
						FriendID: TO_USER_ID,
					}).
					Return(&models.UserFriend{
						UserID:   FROM_USER_ID,
						FriendID: TO_USER_ID,
					}, nil).
					Once()

				bxMock.EXPECT().
					SaveFriendRequestUpdatedID(mockTxCtx, REQ_ID).
					Return(nil).
					Once()

				return uc.Deps{
					RepoFriendReq: frReqMock,
					RepoFriend:    frMock,
					RepoOutbox: bxMock,
					TxMan:         txManMock,
				}
			},
			want: &models.FriendRequest{
				Status:     models.FriendRequestStatusAccepted,
				FromUserID: FROM_USER_ID,
				ToUserID:   TO_USER_ID,
			},
			assertErr: assert.NoError,
		},
		{
			name: "Test 2. Negative: RepoFriendReq.FetchById error",
			args: args{
				reqId: dto.FriendRequestID(REQ_ID),
			},
			mock: func(t *testing.T) uc.Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)
				frMock := mocks.NewFriendRepository(t)
				txManMock := mocks.NewTxManager(t)

				frReqMock.EXPECT().
					FetchById(mockCtx, REQ_ID).
					Return(nil, errors.New("RepoFriendReq.FetchById error")).
					Once()

				return uc.Deps{
					RepoFriendReq: frReqMock,
					RepoFriend:    frMock,
					TxMan:         txManMock,
				}
			},
			want:      nil,
			assertErr: assert.Error,
		},
		{
			name: "Test 3. Negative: Already Accepted Request",
			args: args{
				reqId: dto.FriendRequestID(REQ_ID),
			},
			mock: func(t *testing.T) uc.Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)
				frMock := mocks.NewFriendRepository(t)
				txManMock := mocks.NewTxManager(t)

				frReqMock.EXPECT().
					FetchById(mockCtx, REQ_ID).
					Return(&models.FriendRequest{
						ID:         REQ_ID,
						Status:     models.FriendRequestStatusAccepted,
						FromUserID: FROM_USER_ID,
						ToUserID:   TO_USER_ID,
					}, nil).
					Once()

				return uc.Deps{
					RepoFriendReq: frReqMock,
					RepoFriend:    frMock,
					TxMan:         txManMock,
				}
			},
			want:      nil,
			assertErr: assert.Error,
		},
		{
			name: "Test 4. Negative: Already Declined Request",
			args: args{
				reqId: dto.FriendRequestID(REQ_ID),
			},
			mock: func(t *testing.T) uc.Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)
				frMock := mocks.NewFriendRepository(t)
				txManMock := mocks.NewTxManager(t)

				frReqMock.EXPECT().
					FetchById(mockCtx, REQ_ID).
					Return(&models.FriendRequest{
						ID:         REQ_ID,
						Status:     models.FriendRequestStatusDeclined,
						FromUserID: FROM_USER_ID,
						ToUserID:   TO_USER_ID,
					}, nil).
					Once()

				return uc.Deps{
					RepoFriendReq: frReqMock,
					RepoFriend:    frMock,
					TxMan:         txManMock,
				}
			},
			want:      nil,
			assertErr: assert.Error,
		},
		{
			name: "Test 5. Negative: Transaction operation error",
			args: args{
				reqId: dto.FriendRequestID(REQ_ID),
			},
			mock: func(t *testing.T) uc.Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)
				frMock := mocks.NewFriendRepository(t)
				txManMock := mocks.NewTxManager(t)

				frReqMock.EXPECT().
					FetchById(mockCtx, REQ_ID).
					Return(&models.FriendRequest{
						ID:         REQ_ID,
						Status:     models.FriendRequestStatusPending,
						FromUserID: FROM_USER_ID,
						ToUserID:   TO_USER_ID,
					}, nil).
					Once()

				// TxMan возвращает ошибку, которая имитирует ошибку внутри транзакции
				txManMock.EXPECT().
					RunReadCommitted(mockCtx, mock.AnythingOfType("func(context.Context) error")).
					Return(errors.New("AcceptFriendRequest: any transaction error")).
					Once()

				return uc.Deps{
					RepoFriendReq: frReqMock,
					RepoFriend:    frMock,
					TxMan:         txManMock,
				}
			},
			want:      nil,
			assertErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// ARRANGE
			deps := tt.mock(t)       // создаем мок объекты
			uc := &uc.SocialUsecase{ // инициализируем Usecase
				Deps: deps,
			}

			// ACT
			got, err := uc.AcceptFriendRequest(mockCtx, tt.args.reqId)

			// ASSERT
			tt.assertErr(t, err)

			if got != nil {
				assert.NotEmpty(t, got.ID)
				assert.Equal(t, tt.want.FromUserID, got.FromUserID)
				assert.Equal(t, tt.want.ToUserID, got.ToUserID)
				assert.Equal(t, tt.want.Status, got.Status)
			}

			// assert.Equal(t, tt.want, got)  // mark test as failed but continue execution
			// require.Equal(t, tt.want, got) // mark test as failed and exit
		})
	}
}
