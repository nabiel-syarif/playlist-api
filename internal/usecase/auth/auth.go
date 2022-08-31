package auth

import (
	"context"

	"github.com/jackc/pgx/v4"
	authModel "github.com/nabiel-syarif/playlist-api/internal/model/auth"
	authRepo "github.com/nabiel-syarif/playlist-api/internal/repo/auth"
	authErr "github.com/nabiel-syarif/playlist-api/pkg/error/auth"
	"github.com/nabiel-syarif/playlist-api/pkg/jwt"
	utils "github.com/nabiel-syarif/playlist-api/pkg/utils"
)

//go:generate moq -out auth_mock_test.go . Repo
type Repo interface {
	GetUserByEmail(ctx context.Context, email string) (authModel.User, error)
	GetUsersByUsersId(ctx context.Context, userIds []int) ([]authModel.User, error)
	GetUserByUserId(ctx context.Context, userId int) (authModel.User, error)
	Register(ctx context.Context, newUser authModel.UserRegistration) (authModel.User, error)
}

func New(repo authRepo.Repo, jwtHelper jwt.JwtHelper) Usecase {
	return &usecase{
		repo:      repo,
		jwtHelper: jwtHelper,
	}
}

var _ Usecase = &usecase{}

func (uc *usecase) Login(ctx context.Context, email, password string) (authModel.JwtLoginData, error) {
	user, err := uc.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return authModel.JwtLoginData{}, authErr.ErrBadCredentials
		}
		return authModel.JwtLoginData{}, err
	}

	isPasswordCorrect := utils.ComparePassword(password, user.Password)
	if !isPasswordCorrect {
		return authModel.JwtLoginData{}, authErr.ErrBadCredentials
	}

	token, err := uc.jwtHelper.TokenFromUser(user.Id)
	if err != nil {
		return authModel.JwtLoginData{}, err
	}

	return authModel.JwtLoginData{
		Token: token,
	}, nil
}

func (uc *usecase) Register(ctx context.Context, user authModel.UserRegistration) (authModel.UserPublic, error) {
	var userPublic authModel.UserPublic
	hashedPass, err := utils.HashPassword(user.Password)
	if err != nil {
		return userPublic, err
	}

	user.Password = hashedPass
	newUser, err := uc.repo.Register(ctx, user)
	if err != nil {
		return userPublic, err
	}

	userPublic.Name = newUser.Name
	userPublic.Email = newUser.Email
	userPublic.Id = newUser.Id

	return userPublic, nil

}
