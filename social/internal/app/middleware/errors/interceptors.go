package errors

import (
	"context"
	"errors"
	"log"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases"
)

// ErrorsUnaryInterceptor - convert any arror to rpc error
func ErrorsUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		//
		if _, ok := status.FromError(err); ok {
			return
		}

		switch {
		case errors.Is(err, models.ErrValidationFailed):
			err = status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, models.ErrInvalidFriendRequestStatus):
			err = status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, usecases.ErrUserNoFriendRequestsIn):
			err = status.Error(codes.NotFound, err.Error())
		case errors.Is(err, usecases.ErrUserNoFriendRequestsOut):
			err = status.Error(codes.NotFound, err.Error())
		case errors.Is(err, usecases.ErrUserNoFriends):
			err = status.Error(codes.NotFound, err.Error())
		case errors.Is(err, usecases.ErrNoFriend):
			err = status.Error(codes.NotFound, err.Error())
		case errors.Is(err, usecases.ErrNoFriendRequest):
			err = status.Error(codes.NotFound, err.Error())
		default:
			err = status.Error(codes.Unknown, err.Error())
		}

		return
	}
}

// RecoveryUnaryInterceptor - перехватывает паники и конвертирует их в gRPC ошибки
func RecoveryUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic recovered: %v\n%s", r, debug.Stack())
				err = status.Error(codes.Internal, "internal server error")
			}
		}()
		resp, err = handler(ctx, req)
		return
	}
}
