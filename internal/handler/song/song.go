package song

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	modelSong "github.com/nabiel-syarif/playlist-api/internal/model/song"
	songUc "github.com/nabiel-syarif/playlist-api/internal/usecase/song"
	"github.com/nabiel-syarif/playlist-api/pkg/response"
	"github.com/nabiel-syarif/playlist-api/pkg/utils"
)

type Handler interface {
	SaveSong(w http.ResponseWriter, r *http.Request)
	UpdateSong(w http.ResponseWriter, r *http.Request)
	DeleteSong(w http.ResponseWriter, r *http.Request)
	ListSong(w http.ResponseWriter, r *http.Request)
	GetSongById(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	usecase songUc.Usecase
}

func New(uc songUc.Usecase) Handler {
	return &handler{
		usecase: uc,
	}
}

func (h *handler) SaveSong(w http.ResponseWriter, r *http.Request) {
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

	duration, err := strconv.ParseFloat(r.FormValue("duration"), 64)
	if err != nil {
		resp.Status = "FAILED"
		resp.Error = "Duration should be a number"
		return
	}

	insertSongReq := modelSong.InsertSongRequest{
		Title:     r.FormValue("title"),
		Performer: r.FormValue("performer"),
		Genre:     r.FormValue("genre"),
		Duration:  duration,
	}

	newSong, err := h.usecase.SaveSong(r.Context(), insertSongReq)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = newSong
		resp.Status = "SUCCESS"
		status = http.StatusCreated
	}
}

func (h *handler) UpdateSong(w http.ResponseWriter, r *http.Request) {
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

	songId, err := strconv.Atoi(chi.URLParam(r, "songId"))
	if err != nil {
		resp.Error = "Song id should be a number"
		resp.Status = "FAILED"
		return
	}

	duration, err := strconv.ParseFloat(r.FormValue("duration"), 64)
	if err != nil {
		resp.Status = "FAILED"
		resp.Error = "Duration should be a number"
		return
	}

	updateSongReq := modelSong.InsertSongRequest{
		Title:     r.FormValue("title"),
		Performer: r.FormValue("performer"),
		Genre:     r.FormValue("genre"),
		Duration:  duration,
	}

	song, err := h.usecase.UpdateSong(r.Context(), songId, updateSongReq)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = song
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) DeleteSong(w http.ResponseWriter, r *http.Request) {
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

	songId, err := strconv.Atoi(chi.URLParam(r, "songId"))
	if err != nil {
		resp.Error = "Song id should be a number"
		resp.Status = "FAILED"
		return
	}

	err = h.usecase.DeleteSong(r.Context(), songId)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) ListSong(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.usecase.GetSongs(r.Context())
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = data
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) GetSongById(w http.ResponseWriter, r *http.Request) {
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

	songId, err := strconv.Atoi(chi.URLParam(r, "songId"))
	if err != nil {
		resp.Error = "Song id should be a number"
		resp.Status = "FAILED"
		return
	}

	song, err := h.usecase.GetSongById(r.Context(), songId)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = song
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}
