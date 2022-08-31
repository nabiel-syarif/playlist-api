package playlist

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	playlistModel "github.com/nabiel-syarif/playlist-api/internal/model/playlist"
	playlistUc "github.com/nabiel-syarif/playlist-api/internal/usecase/playlist"
	"github.com/nabiel-syarif/playlist-api/pkg/response"
	"github.com/nabiel-syarif/playlist-api/pkg/utils"
)

//go:generate moq -out playlist_mock_test.go . Usecase
type Usecase interface {
	SavePlaylist(ctx context.Context, req playlistModel.AddPlaylistRequest) (playlistModel.PlaylistAggregated, error)
	AttachSongToPlaylist(ctx context.Context, userId, playlistId, songId int) error
	DetachSongFromPlaylist(ctx context.Context, userId, playlistId, songId int) error
	ListPlaylists(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error)
	GetPlaylistById(ctx context.Context, userId, playlistId int) (playlistModel.PlaylistAggregated, error)
	UpdatePlaylist(ctx context.Context, req playlistModel.UpdatePlaylistRequest) (playlistModel.PlaylistAggregated, error)
	DeletePlaylist(ctx context.Context, userId, playlistId int) error
}

type Handler interface {
	ListPlaylist(w http.ResponseWriter, r *http.Request)
	GetPlaylistById(w http.ResponseWriter, r *http.Request)
	SavePlaylist(w http.ResponseWriter, r *http.Request)
	UpdatePlaylist(w http.ResponseWriter, r *http.Request)
	DeletePlaylist(w http.ResponseWriter, r *http.Request)
	AttachSongToPlaylist(w http.ResponseWriter, r *http.Request)
	DetachSongFromPlaylist(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	usecase playlistUc.Usecase
}

func New(uc playlistUc.Usecase) Handler {
	return &handler{
		usecase: uc,
	}
}

func (h *handler) ListPlaylist(w http.ResponseWriter, r *http.Request) {
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

	data, err := h.usecase.ListPlaylists(r.Context(), userId)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = data
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) SavePlaylist(w http.ResponseWriter, r *http.Request) {
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

	playlist, err := h.usecase.SavePlaylist(r.Context(), playlistModel.AddPlaylistRequest{
		Name:  r.FormValue("name"),
		Owner: userId,
	})
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = playlist
		resp.Status = "SUCCESS"
		status = http.StatusCreated
	}
}

func (h *handler) GetPlaylistById(w http.ResponseWriter, r *http.Request) {
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
		resp.Error = "Failed to get user id from context"
		resp.Status = "FAILED"
		return
	}

	playlistId, err := strconv.Atoi(chi.URLParam(r, "playlistId"))
	if err != nil {
		resp.Error = "Playlist id should be a number"
		resp.Status = "FAILED"
		return
	}

	playlist, err := h.usecase.GetPlaylistById(r.Context(), userId, playlistId)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = playlist
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) AttachSongToPlaylist(w http.ResponseWriter, r *http.Request) {
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

	playlistIdStr := r.FormValue("playlist_id")
	songIdStr := r.FormValue("song_id")
	playlistId, err := strconv.Atoi(playlistIdStr)
	if err != nil {
		resp.Error = "Playlist id should be a number"
		resp.Status = "FAILED"
		return
	}
	songId, err := strconv.Atoi(songIdStr)
	if err != nil {
		resp.Error = "Song id should be a number"
		resp.Status = "FAILED"
		return
	}

	userId, err := utils.GetUserIdFromContext(r.Context())
	if err != nil {
		resp.Error = "Failed to get user id from context"
		resp.Status = "FAILED"
		return
	}

	err = h.usecase.AttachSongToPlaylist(r.Context(), userId, playlistId, songId)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) DetachSongFromPlaylist(w http.ResponseWriter, r *http.Request) {
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

	playlistIdStr := r.FormValue("playlist_id")
	songIdStr := r.FormValue("song_id")
	playlistId, err := strconv.Atoi(playlistIdStr)
	if err != nil {
		resp.Error = "Playlist id should be a number"
		resp.Status = "FAILED"
		return
	}
	songId, err := strconv.Atoi(songIdStr)
	if err != nil {
		resp.Error = "Song id should be a number"
		resp.Status = "FAILED"
		return
	}

	userId, err := utils.GetUserIdFromContext(r.Context())
	if err != nil {
		resp.Error = "Failed to get user id from context"
		resp.Status = "FAILED"
		return
	}

	err = h.usecase.DetachSongFromPlaylist(r.Context(), userId, playlistId, songId)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) UpdatePlaylist(w http.ResponseWriter, r *http.Request) {
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
		resp.Error = "Failed to get user id from context"
		resp.Status = "FAILED"
		return
	}

	playlistId, err := strconv.Atoi(chi.URLParam(r, "playlistId"))
	if err != nil {
		resp.Error = "playlist id should be a number"
		resp.Status = "FAILED"
		return
	}

	data, err := h.usecase.UpdatePlaylist(r.Context(), playlistModel.UpdatePlaylistRequest{
		Id:    playlistId,
		Name:  r.FormValue("name"),
		Owner: userId,
	})
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Data = data
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}

func (h *handler) DeletePlaylist(w http.ResponseWriter, r *http.Request) {
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
		resp.Error = "Failed to get user id from context"
		resp.Status = "FAILED"
		return
	}

	playlistId, err := strconv.Atoi(chi.URLParam(r, "playlistId"))
	if err != nil {
		resp.Error = "playlist id should be a number"
		resp.Status = "FAILED"
		return
	}

	err = h.usecase.DeletePlaylist(r.Context(), userId, playlistId)
	if err != nil {
		resp.Error = err.Error()
		resp.Status = "FAILED"
	} else {
		resp.Status = "SUCCESS"
		status = http.StatusOK
	}
}
