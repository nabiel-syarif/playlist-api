package playlist

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	playlistModel "github.com/nabiel-syarif/playlist-api/internal/model/playlist"
	songModel "github.com/nabiel-syarif/playlist-api/internal/model/song"
	playlistErr "github.com/nabiel-syarif/playlist-api/pkg/error/playlist"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func getPlaylistAggregatedForTest(p playlistModel.Playlist) playlistModel.PlaylistAggregated {
	var playlistCopy playlistModel.Playlist
	if p.Name != "" {
		playlistCopy.Name = p.Name
	} else {
		playlistCopy.Name = "Playlist"
	}
	if p.Id != 0 {
		playlistCopy.Id = 0
	} else {
		playlistCopy.Id = 1
	}
	if p.OwnerId != 0 {
		playlistCopy.OwnerId = p.OwnerId
	} else {
		playlistCopy.OwnerId = 1
	}

	return playlistModel.PlaylistAggregated{
		Id:        playlistCopy.Id,
		Name:      playlistCopy.Name,
		Songs:     make([]songModel.Song, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestNew(t *testing.T) {
	repo := &RepoMock{}
	uc := New(repo)
	require.NotEmpty(t, uc)
}

func TestUsecase_SavePlaylist(t *testing.T) {
	repo := &RepoMock{
		SavePlaylistFunc: func(ctx context.Context, p playlistModel.AddPlaylistRequest) (playlistModel.PlaylistAggregated, error) {
			assert.Equal(t, p.Name, "Playlist")
			assert.Equal(t, p.Owner, 1)
			return getPlaylistAggregatedForTest(playlistModel.Playlist{}), nil
		},
	}

	uc := New(repo)
	p, err := uc.SavePlaylist(context.Background(), playlistModel.AddPlaylistRequest{
		Name:  "Playlist",
		Owner: 1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, p.Id, 1)
}

func TestUsecase_ListPlaylists(t *testing.T) {
	type args struct {
		getUc   func() Usecase
		wantErr bool
	}
	testCases := []struct {
		desc     string
		args     args
		callback func([]playlistModel.PlaylistAggregated, error)
	}{
		{
			desc: "case 1 -> success list all playlist for specific user",
			args: args{
				wantErr: false,
				getUc: func() Usecase {
					return New(&RepoMock{
						ListPlaylistsFunc: func(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error) {
							count := 5
							playlists := make([]playlistModel.PlaylistAggregated, 0)
							for i := 0; i < count; i++ {
								playlists = append(playlists, getPlaylistAggregatedForTest(playlistModel.Playlist{
									Id:      i,
									Name:    fmt.Sprintf("Playlist %d", i),
									OwnerId: i,
								}))
							}

							return playlists, nil
						},
					})
				},
			},
			callback: func(pa []playlistModel.PlaylistAggregated, err error) {
				assert.Len(t, pa, 5)
			},
		},
		{
			desc: "case 2 -> error when trying to list playlists",
			args: args{
				wantErr: true,
				getUc: func() Usecase {
					return New(&RepoMock{
						ListPlaylistsFunc: func(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error) {
							return nil, errors.New("Failed to get playlists")
						},
					})
				},
			},
			callback: func(pa []playlistModel.PlaylistAggregated, err error) {},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p, err := tC.args.getUc().ListPlaylists(context.Background(), 1)
			if (err != nil) != tC.args.wantErr {
				t.Fatalf("got err : %v, want error : %v", err, tC.args.wantErr)
			}
			tC.callback(p, err)
		})
	}
}

func TestUsecase_GetPlaylistById(t *testing.T) {
	type args struct {
		getUc   func() Usecase
		wantErr bool
	}
	userId, playlistId := 1, 1
	testCases := []struct {
		desc     string
		args     args
		callback func(playlistModel.PlaylistAggregated, error)
	}{
		{
			desc: "case 1 -> success get playlist by id",
			args: args{
				wantErr: false,
				getUc: func() Usecase {
					return New(&RepoMock{
						GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
							assert.Equal(t, 1, playlistId)
							return playlistModel.Playlist{
								Id:        1,
								Name:      "Playlist",
								OwnerId:   1,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}, nil
						},
						GetSongsFromPlaylistIdFunc: func(ctx context.Context, playlistId int) ([]songModel.Song, error) {
							return make([]songModel.Song, 0), nil
						},
					})
				},
			},
			callback: func(p playlistModel.PlaylistAggregated, err error) {
				assert.NotEmpty(t, p)
				assert.Equal(t, p.Id, 1)
				assert.Equal(t, p.Name, "Playlist")
			},
		},
		{
			desc: "case 2 -> error not found when trying to get playlist by id",
			args: args{
				wantErr: true,
				getUc: func() Usecase {
					return New(&RepoMock{
						GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
							return playlistModel.Playlist{}, pgx.ErrNoRows
						},
						GetSongsFromPlaylistIdFunc: func(ctx context.Context, playlistId int) ([]songModel.Song, error) {
							return make([]songModel.Song, 0), nil
						},
					})
				},
			},
			callback: func(p playlistModel.PlaylistAggregated, err error) {
				assert.ErrorIs(t, err, playlistErr.ErrPlaylistNotFound)
			},
		},
		{
			desc: "case 3 -> error access (not the owner) when trying to get playlist by id",
			args: args{
				wantErr: true,
				getUc: func() Usecase {
					return New(&RepoMock{
						GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
							return playlistModel.Playlist{
								Id:        1,
								Name:      "Playlist",
								OwnerId:   userId + 1,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}, nil
						},
						GetSongsFromPlaylistIdFunc: func(ctx context.Context, playlistId int) ([]songModel.Song, error) {
							return make([]songModel.Song, 0), nil
						},
					})
				},
			},
			callback: func(p playlistModel.PlaylistAggregated, err error) {
				assert.ErrorIs(t, err, playlistErr.ErrNotPlaylistOwner)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p, err := tC.args.getUc().GetPlaylistById(context.Background(), userId, playlistId)
			if (err != nil) != tC.args.wantErr {
				t.Fatalf("got err : %v, want error : %v", err, tC.args.wantErr)
			}
			tC.callback(p, err)
		})
	}
}

func TestUsecase_UpdatePlaylist(t *testing.T) {
	type args struct {
		getUc   func() Usecase
		wantErr bool
	}
	testCases := []struct {
		desc     string
		args     args
		callback func(playlistModel.PlaylistAggregated, error)
	}{
		{
			desc: "case 1 -> success update playlist",
			args: args{
				wantErr: false,
				getUc: func() Usecase {
					return New(&RepoMock{
						UpdatePlaylistFunc: func(ctx context.Context, p playlistModel.UpdatePlaylistRequest) (playlistModel.Playlist, error) {
							require.Equal(t, p.Id, 1)
							require.Equal(t, p.Name, "Playlist Updated")
							return playlistModel.Playlist{
								Id:        1,
								Name:      "Playlist Updated",
								OwnerId:   1,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}, nil
						},
						GetSongsFromPlaylistIdFunc: func(ctx context.Context, playlistId int) ([]songModel.Song, error) {
							return make([]songModel.Song, 0), nil
						},
						GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
							require.Equal(t, playlistId, 1)

							return playlistModel.Playlist{
								Id:        1,
								Name:      "Playlist Updated",
								OwnerId:   1,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}, nil
						},
					})
				},
			},
			callback: func(p playlistModel.PlaylistAggregated, err error) {
				require.Equal(t, p.Id, 1)
				require.Equal(t, p.Name, "Playlist Updated")
				require.NotNil(t, p.Songs)
				require.Len(t, p.Songs, 0)
			},
		},
		{
			desc: "case 2 -> error not found when trying to update playlist",
			args: args{
				wantErr: true,
				getUc: func() Usecase {
					return New(&RepoMock{
						UpdatePlaylistFunc: func(ctx context.Context, p playlistModel.UpdatePlaylistRequest) (playlistModel.Playlist, error) {
							return playlistModel.Playlist{}, pgx.ErrNoRows
						},
						GetSongsFromPlaylistIdFunc: func(ctx context.Context, playlistId int) ([]songModel.Song, error) {
							return make([]songModel.Song, 0), nil
						},
						GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
							require.Equal(t, playlistId, 1)

							return playlistModel.Playlist{
								Id:        1,
								Name:      "Playlist Updated",
								OwnerId:   1,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}, nil
						},
					})
				},
			},
			callback: func(p playlistModel.PlaylistAggregated, err error) {
				assert.ErrorIs(t, err, playlistErr.ErrPlaylistNotFound)
			},
		},
		{
			desc: "case 3 -> error access (not the owner) when trying to update playlist",
			args: args{
				wantErr: true,
				getUc: func() Usecase {
					return New(&RepoMock{
						UpdatePlaylistFunc: func(ctx context.Context, p playlistModel.UpdatePlaylistRequest) (playlistModel.Playlist, error) {
							return playlistModel.Playlist{}, pgx.ErrNoRows
						},
						GetSongsFromPlaylistIdFunc: func(ctx context.Context, playlistId int) ([]songModel.Song, error) {
							return make([]songModel.Song, 0), nil
						},
						GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
							require.Equal(t, playlistId, 1)

							return playlistModel.Playlist{
								Id:        1,
								Name:      "Playlist Updated",
								OwnerId:   2,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}, nil
						},
					})
				},
			},
			callback: func(p playlistModel.PlaylistAggregated, err error) {
				assert.ErrorIs(t, err, playlistErr.ErrNotPlaylistOwner)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p, err := tC.args.getUc().UpdatePlaylist(context.Background(), playlistModel.UpdatePlaylistRequest{
				Id:    1,
				Name:  "Playlist Updated",
				Owner: 1,
			})
			if (err != nil) != tC.args.wantErr {
				t.Fatalf("got err : %v, want error : %v", err, tC.args.wantErr)
			}
			tC.callback(p, err)
		})
	}
}

func TestUsecase_DeletePlaylist(t *testing.T) {
	type args struct {
		getUc   func() Usecase
		wantErr bool
	}
	ownerId := 1
	testCases := []struct {
		desc string
		args args
	}{
		{
			desc: "case 1 -> success delete playlist",
			args: args{
				wantErr: false,
				getUc: func() Usecase {
					return New(&RepoMock{
						GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
							return playlistModel.Playlist{
								Id:        playlistId,
								Name:      "Playlist",
								OwnerId:   ownerId,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}, nil
						},
						DeletePlaylistFunc: func(ctx context.Context, playlistId int) error {
							return nil
						},
					})
				},
			},
		},
		{
			desc: "case 2 -> error not found when trying to delete playlist",
			args: args{
				wantErr: true,
				getUc: func() Usecase {
					return New(&RepoMock{
						GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
							return playlistModel.Playlist{
								Id:        playlistId,
								Name:      "Playlist",
								OwnerId:   ownerId,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}, nil
						},
						DeletePlaylistFunc: func(ctx context.Context, playlistId int) error {
							return pgx.ErrNoRows
						},
					})
				},
			},
		},
		{
			desc: "case 3 -> error access when trying to delete playlist",
			args: args{
				wantErr: true,
				getUc: func() Usecase {
					return New(&RepoMock{
						GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
							return playlistModel.Playlist{
								Id:        playlistId,
								Name:      "Playlist",
								OwnerId:   ownerId + 1,
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}, nil
						},
						DeletePlaylistFunc: func(ctx context.Context, playlistId int) error {
							return nil
						},
					})
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			err := tC.args.getUc().DeletePlaylist(context.Background(), 1, 1)
			if (err != nil) != tC.args.wantErr {
				t.Fatalf("got err : %v, want error : %v", err, tC.args.wantErr)
			}
		})
	}
}

func TestUsecase_CheckPlaylistAccess(t *testing.T) {
	ownerId, playlistId := 1, 1
	type args struct {
		getUc   func() usecase
		wantErr bool
	}
	testCases := []struct {
		desc     string
		args     args
		callback func(error)
	}{
		{
			desc: "case 1 -> success check playlist access",
			args: args{
				wantErr: false,
				getUc: func() usecase {
					return usecase{
						repo: &RepoMock{
							GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
								return playlistModel.Playlist{
									Id:        playlistId,
									Name:      "Playlist",
									OwnerId:   ownerId,
									CreatedAt: time.Now(),
									UpdatedAt: time.Now(),
								}, nil
							},
						}}
				},
			},
			callback: func(err error) {
				require.NoError(t, err)
			},
		},
		{
			desc: "case 2 -> error no access on playlist",
			args: args{
				wantErr: true,
				getUc: func() usecase {
					return usecase{
						repo: &RepoMock{
							GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
								return playlistModel.Playlist{
									Id:        playlistId,
									Name:      "Playlist",
									OwnerId:   ownerId + 1,
									CreatedAt: time.Now(),
									UpdatedAt: time.Now(),
								}, nil
							},
						}}
				},
			},
			callback: func(err error) {
				require.ErrorIs(t, err, playlistErr.ErrNotPlaylistOwner)
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			uc := tC.args.getUc()
			err := uc.CheckPlaylistAccess(context.Background(), ownerId, playlistId)
			if (err != nil) != tC.args.wantErr {
				t.Fatalf("got err : %v, want error : %v", err, tC.args.wantErr)
			}
		})
	}
}
