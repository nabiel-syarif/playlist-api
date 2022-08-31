package auth

import (
	"context"
	"log"
	"net/http"

	authModel "github.com/nabiel-syarif/playlist-api/internal/model/auth"
	authUc "github.com/nabiel-syarif/playlist-api/internal/usecase/auth"
	"github.com/nabiel-syarif/playlist-api/pkg/response"
	"github.com/nabiel-syarif/playlist-api/pkg/utils"
)

type Handler interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	usecase authUc.Usecase
}

//go:generate moq -out auth_mock_test.go . Usecase
type Usecase interface {
	Login(ctx context.Context, email, password string) (authModel.JwtLoginData, error)
	Register(ctx context.Context, user authModel.UserRegistration) (authModel.UserPublic, error)
}

func New(uc authUc.Usecase) Handler {
	return &handler{
		usecase: uc,
	}
}

func (h *handler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		resp   response.StandardResponse
		status = http.StatusBadRequest
	)
	defer func() {
		err := utils.ResponseWithJson(w, status, resp)
		if err != nil {
			log.Printf("Failed to write response, err : %v\n", err)
		}
	}()

	user := authModel.UserRegistration{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	data, err := h.usecase.Register(r.Context(), user)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = data
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		resp   response.StandardResponse
		status = http.StatusBadRequest
	)
	defer func() {
		err := utils.ResponseWithJson(w, status, resp)
		if err != nil {
			log.Printf("Failed to write response, err : %v\n", err)
		}
	}()

	email := r.FormValue("email")
	password := r.FormValue("password")

	data, err := h.usecase.Login(r.Context(), email, password)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = data
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}
