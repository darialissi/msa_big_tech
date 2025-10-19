package errors

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/darialissi/msa_big_tech/users/internal/app/models"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases"
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
		case errors.Is(err, usecases.ErrExistedNickname):
			err = status.Error(codes.AlreadyExists, err.Error())
		case errors.Is(err, usecases.ErrNoProfileFound):
			err = status.Error(codes.NotFound, err.Error())
		default:
			err = status.Error(codes.Unknown, err.Error())
		}

		return
	}
}
