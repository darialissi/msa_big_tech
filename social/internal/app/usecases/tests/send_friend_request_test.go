package tests

import (
	"testing"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	uc "github.com/darialissi/msa_big_tech/social/internal/app/usecases"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/tests/mocks"
)

func Test_SendFriendRequest_whitebox_mockery(t *testing.T) {
	t.Parallel()

	mockCtx := &mocks.MockContext{}
	mockTxCtx := &mocks.MockContext{}

	REQ_ID := models.FriendRequestID(uuid.New().String())
	FROM_USER_ID := models.UserID(uuid.New().String())
	TO_USER_ID := models.UserID(uuid.New().String())

	// ARRANGE
	type args struct {
		frReqInfo *dto.SendFriendRequest
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
				frReqInfo: &dto.SendFriendRequest{
					FromUserID: dto.UserID(FROM_USER_ID),
					ToUserID:   dto.UserID(TO_USER_ID),
				},
			},
			mock: func(t *testing.T) uc.Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)
				frMock := mocks.NewFriendRepository(t)
				bxMock := mocks.NewOutboxRepository(t)
				txManMock := mocks.NewTxManager(t)

				txManMock.EXPECT().
					RunReadCommitted(mockCtx, mock.AnythingOfType("func(context.Context) error")).
					Run(func(ctx context.Context, fn func(context.Context) error) {
						// Имитируем выполнение внутри транзакции
						fn(mockTxCtx)
					}).
					Return(nil).
					Once()

				frReqMock.EXPECT().
					Save(mockTxCtx, &models.FriendRequest{
						Status:     models.FriendRequestStatusPending,
						FromUserID: FROM_USER_ID,
						ToUserID:   TO_USER_ID,
					}).
					Return(&models.FriendRequest{
						ID:         REQ_ID,
						Status:     models.FriendRequestStatusPending,
						FromUserID: FROM_USER_ID,
						ToUserID:   TO_USER_ID,
					}, nil).
					Once()

				bxMock.EXPECT().
					SaveFriendRequestCreatedID(mockTxCtx, REQ_ID).
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
				Status:     models.FriendRequestStatusPending,
				FromUserID: models.UserID(FROM_USER_ID),
				ToUserID:   models.UserID(TO_USER_ID),
			},
			assertErr: assert.NoError,
		},
		{
			name: "Test 2. Negative: Request to yourself",
			args: args{
				frReqInfo: &dto.SendFriendRequest{
					FromUserID: dto.UserID(FROM_USER_ID),
					ToUserID:   dto.UserID(FROM_USER_ID),
				},
			},
			mock: func(t *testing.T) uc.Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)
				frMock := mocks.NewFriendRepository(t)
				bxMock := mocks.NewOutboxRepository(t)
				txManMock := mocks.NewTxManager(t)

				return uc.Deps{
					RepoFriendReq: frReqMock,
					RepoFriend:    frMock,
					RepoOutbox: bxMock,
					TxMan:         txManMock,
				}
			},
			want:      nil,
			assertErr: assert.Error,
		},
		{
			name: "Test 3. Negative: Transaction operation error",
			args: args{
				frReqInfo: &dto.SendFriendRequest{
					FromUserID: dto.UserID(FROM_USER_ID),
					ToUserID:   dto.UserID(TO_USER_ID),
				},
			},
			mock: func(t *testing.T) uc.Deps {
				frReqMock := mocks.NewFriendRequestRepository(t)
				frMock := mocks.NewFriendRepository(t)
				bxMock := mocks.NewOutboxRepository(t)
				txManMock := mocks.NewTxManager(t)

				// TxMan возвращает ошибку, которая имитирует ошибку внутри транзакции
				txManMock.EXPECT().
					RunReadCommitted(mockCtx, mock.AnythingOfType("func(context.Context) error")).
					Return(errors.New("AcceptFriendRequest: any transaction error")).
					Once()

				return uc.Deps{
					RepoFriendReq: frReqMock,
					RepoFriend:    frMock,
					RepoOutbox: bxMock,
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
			got, err := uc.SendFriendRequest(mockCtx, tt.args.frReqInfo)

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
