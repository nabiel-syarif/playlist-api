package playlist

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	collabModel "github.com/nabiel-syarif/playlist-api/internal/model/collaboration"
	collabUc "github.com/nabiel-syarif/playlist-api/internal/usecase/collaboration"
	"github.com/nabiel-syarif/playlist-api/pkg/response"
	"github.com/nabiel-syarif/playlist-api/pkg/utils"
)

//go:generate moq -out collaboration_mock_test.go . Usecase
type Usecase interface {
	AddCollaborator(w http.ResponseWriter, r *http.Request) error
	RemoveCollaborator(w http.ResponseWriter, r *http.Request) error
}

type Handler interface {
	AddCollaborator(w http.ResponseWriter, r *http.Request)
	RemoveCollaborator(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	usecase collabUc.Usecase
}

func New(uc collabUc.Usecase) Handler {
	return &handler{
		usecase: uc,
	}
}

func (h *handler) AddCollaborator(w http.ResponseWriter, r *http.Request) {
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

	userId, err := utils.GetUserIdFromContext(r.Context())
	if err != nil {
		resp.Error = "Can't get user id from context"
		resp.Status = "FAILED"
		return
	}

	playlistId, err := strconv.Atoi(chi.URLParam(r, "playlistId"))
	if err != nil {
		resp.Error = "Can't get playlist id from url"
		resp.Status = "FAILED"
		return
	}

	requests := make([]collabModel.AddCollaboratorRequest, 0)
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		resp.Status = "FAILED"
		resp.Error = err.Error()
		status = http.StatusInternalServerError
		return
	}

	err = json.Unmarshal(bytes, &requests)
	if err != nil {
		resp.Status = "FAILED"
		resp.Error = err.Error()
		status = http.StatusInternalServerError
		return
	}

	err = h.usecase.AddCollaborator(r.Context(), playlistId, userId, requests)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) RemoveCollaborator(w http.ResponseWriter, r *http.Request) {
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

	userId, err := utils.GetUserIdFromContext(r.Context())
	if err != nil {
		resp.Error = "Can't get user id from context"
		resp.Status = "FAILED"
		return
	}

	playlistId, err := strconv.Atoi(chi.URLParam(r, "playlistId"))
	if err != nil {
		resp.Error = "Can't get playlist id from url"
		resp.Status = "FAILED"
		return
	}

	requests := make([]collabModel.RemoveCollaboratorRequest, 0)
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		resp.Status = "FAILED"
		resp.Error = err.Error()
		status = http.StatusInternalServerError
		return
	}

	err = json.Unmarshal(bytes, &requests)
	if err != nil {
		resp.Status = "FAILED"
		resp.Error = err.Error()
		status = http.StatusInternalServerError
		return
	}

	err = h.usecase.RemoveCollaborator(r.Context(), playlistId, userId, requests)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}
