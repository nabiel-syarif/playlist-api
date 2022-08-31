package playlist

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	playlistModel "github.com/nabiel-syarif/playlist-api/internal/model/playlist"
	"github.com/nabiel-syarif/playlist-api/internal/model/song"
	"github.com/stretchr/testify/require"
)

func TestHandler_ListPlaylists(t *testing.T) {
	type args struct {
		prepareRequest func() *http.Request
	}
	testCases := []struct {
		desc     string
		args     args
		uc       Usecase
		callback func(*testing.T, map[string]interface{})
	}{
		{
			desc: "case 1 -> success list playlists",
			args: args{
				prepareRequest: func() *http.Request {
					// body := &bytes.Buffer{}
					// writer := multipart.NewWriter(body)
					// fw, _ := writer.CreateFormField("name")
					// _, _ = io.Copy(fw, strings.NewReader("Nabiel"))

					// writer.Close()
					req, err := http.NewRequest(http.MethodPost, "/v1/playlists", nil)
					req = req.WithContext(context.WithValue(req.Context(), "userId", 1))
					require.NoError(t, err)
					return req
				},
			},
			uc: &UsecaseMock{
				ListPlaylistsFunc: func(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error) {
					return []playlistModel.PlaylistAggregated{
						{
							Id:        1,
							Name:      "Playlist 1",
							Songs:     []song.Song{},
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
						},
						{
							Id:        2,
							Name:      "Playlist 2",
							Songs:     []song.Song{},
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
						},
					}, nil
				},
			},
			callback: func(t *testing.T, m map[string]interface{}) {
				require.Contains(t, m, "status")
				require.Equal(t, m["status"], "SUCCESS")
				require.Len(t, m["data"], 2)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			h := New(tC.uc)

			router := chi.NewRouter()
			router.Method(http.MethodPost, "/v1/playlists", http.HandlerFunc(h.ListPlaylist))

			recorder := httptest.NewRecorder()
			request := tC.args.prepareRequest()
			router.ServeHTTP(recorder, request)
			var response map[string]interface{}
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			require.NoError(t, err, "unmarshall no error")
			tC.callback(t, response)
		})
	}
}
