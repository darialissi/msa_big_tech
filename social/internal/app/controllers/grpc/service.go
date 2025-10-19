package social_grpc

import (
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases"
	social "github.com/darialissi/msa_big_tech/social/pkg"
)

type service struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	social.UnimplementedSocialServiceServer

	SocialUsecase *usecases.SocialUsecase
}

func NewServer(socialUC *usecases.SocialUsecase) *service {
	return &service{
		SocialUsecase: socialUC,
	}
}
