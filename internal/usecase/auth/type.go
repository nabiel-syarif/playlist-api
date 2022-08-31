package auth

import (
	"context"

	authModel "github.com/nabiel-syarif/playlist-api/internal/model/auth"
	authRepo "github.com/nabiel-syarif/playlist-api/internal/repo/auth"
	"github.com/nabiel-syarif/playlist-api/pkg/jwt"
)

type usecase struct {
	repo      authRepo.Repo
	jwtHelper jwt.JwtHelper
}
type Usecase interface {
	Login(ctx context.Context, email, password string) (authModel.JwtLoginData, error)
	Register(ctx context.Context, user authModel.UserRegistration) (authModel.UserPublic, error)
}

type JwtConfig struct {
	SecretKey string
}
