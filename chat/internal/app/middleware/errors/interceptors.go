package errors

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases"
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
		case errors.Is(err, usecases.ErrNoUserChats):
			err = status.Error(codes.NotFound, err.Error())
		case errors.Is(err, usecases.ErrNoUserMessages):
			err = status.Error(codes.NotFound, err.Error())
		case errors.Is(err, usecases.ErrNotExistedChat):
			err = status.Error(codes.NotFound, err.Error())
		case errors.Is(err, usecases.ErrNoMessages):
			err = status.Error(codes.NotFound, err.Error())
		case errors.Is(err, usecases.ErrNotExistedMessage):
			err = status.Error(codes.NotFound, err.Error())
		default:
			err = status.Error(codes.Unknown, err.Error())
		}

		return
	}
}