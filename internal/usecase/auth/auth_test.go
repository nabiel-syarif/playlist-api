package auth

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/nabiel-syarif/playlist-api/internal/model/auth"
	authErr "github.com/nabiel-syarif/playlist-api/pkg/error/auth"
	"github.com/nabiel-syarif/playlist-api/pkg/jwt"
	"github.com/nabiel-syarif/playlist-api/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	repo := &RepoMock{}
	uc := New(repo, jwt.JwtHelper{})
	require.NotEmpty(t, uc, "Usecase should not empty")
}

func getTestUser() (auth.User, error) {
	hashed, err := utils.HashPassword("rahasia")
	if err != nil {
		return auth.User{}, err
	}
	return auth.User{
		Id:        1,
		Name:      "Nabiel",
		Email:     "nabiel@gmail.com",
		Password:  hashed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now().Add(time.Minute * 5),
	}, nil
}

func TestLogin(t *testing.T) {
	repo := &RepoMock{
		GetUserByEmailFunc: func(ctx context.Context, email string) (auth.User, error) {
			if email == "nabiel@gmail.com" {
				return getTestUser()
			}
			return auth.User{}, pgx.ErrNoRows
		},
	}
	uc := New(repo, jwt.JwtHelper{
		Config: struct {
			SecretKey       string
			TokenExpiration int
		}{
			SecretKey:       "rahasia",
			TokenExpiration: 3600,
		},
	})

	token, err := uc.Login(context.Background(), "nabiel@gmail.com", "rahasia")
	require.NoError(t, err, "Should not return error")
	require.NotEmpty(t, token, "Token should not empty")

	_, err = uc.Login(context.Background(), "nabiel@gmail.com", "wrong password")
	require.Error(t, err)
	require.ErrorIs(t, err, authErr.ErrBadCredentials)

	_, err = uc.Login(context.Background(), "unknown@gmail.com", "wrong password")
	require.Error(t, err, authErr.ErrBadCredentials)
}

func TestRegister(t *testing.T) {
	repo := &RepoMock{
		RegisterFunc: func(ctx context.Context, newUser auth.UserRegistration) (auth.User, error) {
			return getTestUser()
		},
	}

	uc := New(repo, jwt.JwtHelper{
		Config: struct {
			SecretKey       string
			TokenExpiration int
		}{
			SecretKey:       "rahasia",
			TokenExpiration: 3600,
		},
	})
	user, err := uc.Register(context.Background(), auth.UserRegistration{
		Name:     "Nabiel",
		Email:    "nabiel@gmail.com",
		Password: "rahasia",
	})
	require.NoError(t, err, "should not return error")
	require.Len(t, repo.RegisterCalls(), 1)
	require.NotEmpty(t, user, "user should not empty")
}
