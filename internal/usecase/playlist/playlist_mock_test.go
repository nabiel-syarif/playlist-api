// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package playlist

import (
	"context"
	playlistModel "github.com/nabiel-syarif/playlist-api/internal/model/playlist"
	songModel "github.com/nabiel-syarif/playlist-api/internal/model/song"
	"sync"
)

// Ensure, that RepoMock does implement Repo.
// If this is not the case, regenerate this file with moq.
var _ Repo = &RepoMock{}

// RepoMock is a mock implementation of Repo.
//
// 	func TestSomethingThatUsesRepo(t *testing.T) {
//
// 		// make and configure a mocked Repo
// 		mockedRepo := &RepoMock{
// 			AddSongToPlaylistFunc: func(ctx context.Context, playlistId int, songId int) error {
// 				panic("mock out the AddSongToPlaylist method")
// 			},
// 			CheckPlaylistAccessFunc: func(ctx context.Context, playlistId int, userid int) error {
// 				panic("mock out the CheckPlaylistAccess method")
// 			},
// 			DeletePlaylistFunc: func(ctx context.Context, playlistId int) error {
// 				panic("mock out the DeletePlaylist method")
// 			},
// 			GetPlaylistByIdFunc: func(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
// 				panic("mock out the GetPlaylistById method")
// 			},
// 			GetSongsFromPlaylistIdFunc: func(ctx context.Context, playlistId int) ([]songModel.Song, error) {
// 				panic("mock out the GetSongsFromPlaylistId method")
// 			},
// 			ListPlaylistsFunc: func(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error) {
// 				panic("mock out the ListPlaylists method")
// 			},
// 			RemoveSongFromPlaylistFunc: func(ctx context.Context, playlistId int, songId int) error {
// 				panic("mock out the RemoveSongFromPlaylist method")
// 			},
// 			SavePlaylistFunc: func(ctx context.Context, p playlistModel.AddPlaylistRequest) (playlistModel.PlaylistAggregated, error) {
// 				panic("mock out the SavePlaylist method")
// 			},
// 			UpdatePlaylistFunc: func(ctx context.Context, p playlistModel.UpdatePlaylistRequest) (playlistModel.Playlist, error) {
// 				panic("mock out the UpdatePlaylist method")
// 			},
// 		}
//
// 		// use mockedRepo in code that requires Repo
// 		// and then make assertions.
//
// 	}
type RepoMock struct {
	// AddSongToPlaylistFunc mocks the AddSongToPlaylist method.
	AddSongToPlaylistFunc func(ctx context.Context, playlistId int, songId int) error

	// CheckPlaylistAccessFunc mocks the CheckPlaylistAccess method.
	CheckPlaylistAccessFunc func(ctx context.Context, playlistId int, userid int) error

	// DeletePlaylistFunc mocks the DeletePlaylist method.
	DeletePlaylistFunc func(ctx context.Context, playlistId int) error

	// GetPlaylistByIdFunc mocks the GetPlaylistById method.
	GetPlaylistByIdFunc func(ctx context.Context, playlistId int) (playlistModel.Playlist, error)

	// GetSongsFromPlaylistIdFunc mocks the GetSongsFromPlaylistId method.
	GetSongsFromPlaylistIdFunc func(ctx context.Context, playlistId int) ([]songModel.Song, error)

	// ListPlaylistsFunc mocks the ListPlaylists method.
	ListPlaylistsFunc func(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error)

	// RemoveSongFromPlaylistFunc mocks the RemoveSongFromPlaylist method.
	RemoveSongFromPlaylistFunc func(ctx context.Context, playlistId int, songId int) error

	// SavePlaylistFunc mocks the SavePlaylist method.
	SavePlaylistFunc func(ctx context.Context, p playlistModel.AddPlaylistRequest) (playlistModel.PlaylistAggregated, error)

	// UpdatePlaylistFunc mocks the UpdatePlaylist method.
	UpdatePlaylistFunc func(ctx context.Context, p playlistModel.UpdatePlaylistRequest) (playlistModel.Playlist, error)

	// calls tracks calls to the methods.
	calls struct {
		// AddSongToPlaylist holds details about calls to the AddSongToPlaylist method.
		AddSongToPlaylist []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// PlaylistId is the playlistId argument value.
			PlaylistId int
			// SongId is the songId argument value.
			SongId int
		}
		// CheckPlaylistAccess holds details about calls to the CheckPlaylistAccess method.
		CheckPlaylistAccess []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// PlaylistId is the playlistId argument value.
			PlaylistId int
			// Userid is the userid argument value.
			Userid int
		}
		// DeletePlaylist holds details about calls to the DeletePlaylist method.
		DeletePlaylist []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// PlaylistId is the playlistId argument value.
			PlaylistId int
		}
		// GetPlaylistById holds details about calls to the GetPlaylistById method.
		GetPlaylistById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// PlaylistId is the playlistId argument value.
			PlaylistId int
		}
		// GetSongsFromPlaylistId holds details about calls to the GetSongsFromPlaylistId method.
		GetSongsFromPlaylistId []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// PlaylistId is the playlistId argument value.
			PlaylistId int
		}
		// ListPlaylists holds details about calls to the ListPlaylists method.
		ListPlaylists []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UserId is the userId argument value.
			UserId int
		}
		// RemoveSongFromPlaylist holds details about calls to the RemoveSongFromPlaylist method.
		RemoveSongFromPlaylist []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// PlaylistId is the playlistId argument value.
			PlaylistId int
			// SongId is the songId argument value.
			SongId int
		}
		// SavePlaylist holds details about calls to the SavePlaylist method.
		SavePlaylist []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P playlistModel.AddPlaylistRequest
		}
		// UpdatePlaylist holds details about calls to the UpdatePlaylist method.
		UpdatePlaylist []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// P is the p argument value.
			P playlistModel.UpdatePlaylistRequest
		}
	}
	lockAddSongToPlaylist      sync.RWMutex
	lockCheckPlaylistAccess    sync.RWMutex
	lockDeletePlaylist         sync.RWMutex
	lockGetPlaylistById        sync.RWMutex
	lockGetSongsFromPlaylistId sync.RWMutex
	lockListPlaylists          sync.RWMutex
	lockRemoveSongFromPlaylist sync.RWMutex
	lockSavePlaylist           sync.RWMutex
	lockUpdatePlaylist         sync.RWMutex
}

// AddSongToPlaylist calls AddSongToPlaylistFunc.
func (mock *RepoMock) AddSongToPlaylist(ctx context.Context, playlistId int, songId int) error {
	if mock.AddSongToPlaylistFunc == nil {
		panic("RepoMock.AddSongToPlaylistFunc: method is nil but Repo.AddSongToPlaylist was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		PlaylistId int
		SongId     int
	}{
		Ctx:        ctx,
		PlaylistId: playlistId,
		SongId:     songId,
	}
	mock.lockAddSongToPlaylist.Lock()
	mock.calls.AddSongToPlaylist = append(mock.calls.AddSongToPlaylist, callInfo)
	mock.lockAddSongToPlaylist.Unlock()
	return mock.AddSongToPlaylistFunc(ctx, playlistId, songId)
}

// AddSongToPlaylistCalls gets all the calls that were made to AddSongToPlaylist.
// Check the length with:
//     len(mockedRepo.AddSongToPlaylistCalls())
func (mock *RepoMock) AddSongToPlaylistCalls() []struct {
	Ctx        context.Context
	PlaylistId int
	SongId     int
} {
	var calls []struct {
		Ctx        context.Context
		PlaylistId int
		SongId     int
	}
	mock.lockAddSongToPlaylist.RLock()
	calls = mock.calls.AddSongToPlaylist
	mock.lockAddSongToPlaylist.RUnlock()
	return calls
}

// CheckPlaylistAccess calls CheckPlaylistAccessFunc.
func (mock *RepoMock) CheckPlaylistAccess(ctx context.Context, playlistId int, userid int) error {
	if mock.CheckPlaylistAccessFunc == nil {
		panic("RepoMock.CheckPlaylistAccessFunc: method is nil but Repo.CheckPlaylistAccess was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		PlaylistId int
		Userid     int
	}{
		Ctx:        ctx,
		PlaylistId: playlistId,
		Userid:     userid,
	}
	mock.lockCheckPlaylistAccess.Lock()
	mock.calls.CheckPlaylistAccess = append(mock.calls.CheckPlaylistAccess, callInfo)
	mock.lockCheckPlaylistAccess.Unlock()
	return mock.CheckPlaylistAccessFunc(ctx, playlistId, userid)
}

// CheckPlaylistAccessCalls gets all the calls that were made to CheckPlaylistAccess.
// Check the length with:
//     len(mockedRepo.CheckPlaylistAccessCalls())
func (mock *RepoMock) CheckPlaylistAccessCalls() []struct {
	Ctx        context.Context
	PlaylistId int
	Userid     int
} {
	var calls []struct {
		Ctx        context.Context
		PlaylistId int
		Userid     int
	}
	mock.lockCheckPlaylistAccess.RLock()
	calls = mock.calls.CheckPlaylistAccess
	mock.lockCheckPlaylistAccess.RUnlock()
	return calls
}

// DeletePlaylist calls DeletePlaylistFunc.
func (mock *RepoMock) DeletePlaylist(ctx context.Context, playlistId int) error {
	if mock.DeletePlaylistFunc == nil {
		panic("RepoMock.DeletePlaylistFunc: method is nil but Repo.DeletePlaylist was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		PlaylistId int
	}{
		Ctx:        ctx,
		PlaylistId: playlistId,
	}
	mock.lockDeletePlaylist.Lock()
	mock.calls.DeletePlaylist = append(mock.calls.DeletePlaylist, callInfo)
	mock.lockDeletePlaylist.Unlock()
	return mock.DeletePlaylistFunc(ctx, playlistId)
}

// DeletePlaylistCalls gets all the calls that were made to DeletePlaylist.
// Check the length with:
//     len(mockedRepo.DeletePlaylistCalls())
func (mock *RepoMock) DeletePlaylistCalls() []struct {
	Ctx        context.Context
	PlaylistId int
} {
	var calls []struct {
		Ctx        context.Context
		PlaylistId int
	}
	mock.lockDeletePlaylist.RLock()
	calls = mock.calls.DeletePlaylist
	mock.lockDeletePlaylist.RUnlock()
	return calls
}

// GetPlaylistById calls GetPlaylistByIdFunc.
func (mock *RepoMock) GetPlaylistById(ctx context.Context, playlistId int) (playlistModel.Playlist, error) {
	if mock.GetPlaylistByIdFunc == nil {
		panic("RepoMock.GetPlaylistByIdFunc: method is nil but Repo.GetPlaylistById was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		PlaylistId int
	}{
		Ctx:        ctx,
		PlaylistId: playlistId,
	}
	mock.lockGetPlaylistById.Lock()
	mock.calls.GetPlaylistById = append(mock.calls.GetPlaylistById, callInfo)
	mock.lockGetPlaylistById.Unlock()
	return mock.GetPlaylistByIdFunc(ctx, playlistId)
}

// GetPlaylistByIdCalls gets all the calls that were made to GetPlaylistById.
// Check the length with:
//     len(mockedRepo.GetPlaylistByIdCalls())
func (mock *RepoMock) GetPlaylistByIdCalls() []struct {
	Ctx        context.Context
	PlaylistId int
} {
	var calls []struct {
		Ctx        context.Context
		PlaylistId int
	}
	mock.lockGetPlaylistById.RLock()
	calls = mock.calls.GetPlaylistById
	mock.lockGetPlaylistById.RUnlock()
	return calls
}

// GetSongsFromPlaylistId calls GetSongsFromPlaylistIdFunc.
func (mock *RepoMock) GetSongsFromPlaylistId(ctx context.Context, playlistId int) ([]songModel.Song, error) {
	if mock.GetSongsFromPlaylistIdFunc == nil {
		panic("RepoMock.GetSongsFromPlaylistIdFunc: method is nil but Repo.GetSongsFromPlaylistId was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		PlaylistId int
	}{
		Ctx:        ctx,
		PlaylistId: playlistId,
	}
	mock.lockGetSongsFromPlaylistId.Lock()
	mock.calls.GetSongsFromPlaylistId = append(mock.calls.GetSongsFromPlaylistId, callInfo)
	mock.lockGetSongsFromPlaylistId.Unlock()
	return mock.GetSongsFromPlaylistIdFunc(ctx, playlistId)
}

// GetSongsFromPlaylistIdCalls gets all the calls that were made to GetSongsFromPlaylistId.
// Check the length with:
//     len(mockedRepo.GetSongsFromPlaylistIdCalls())
func (mock *RepoMock) GetSongsFromPlaylistIdCalls() []struct {
	Ctx        context.Context
	PlaylistId int
} {
	var calls []struct {
		Ctx        context.Context
		PlaylistId int
	}
	mock.lockGetSongsFromPlaylistId.RLock()
	calls = mock.calls.GetSongsFromPlaylistId
	mock.lockGetSongsFromPlaylistId.RUnlock()
	return calls
}

// ListPlaylists calls ListPlaylistsFunc.
func (mock *RepoMock) ListPlaylists(ctx context.Context, userId int) ([]playlistModel.PlaylistAggregated, error) {
	if mock.ListPlaylistsFunc == nil {
		panic("RepoMock.ListPlaylistsFunc: method is nil but Repo.ListPlaylists was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		UserId int
	}{
		Ctx:    ctx,
		UserId: userId,
	}
	mock.lockListPlaylists.Lock()
	mock.calls.ListPlaylists = append(mock.calls.ListPlaylists, callInfo)
	mock.lockListPlaylists.Unlock()
	return mock.ListPlaylistsFunc(ctx, userId)
}

// ListPlaylistsCalls gets all the calls that were made to ListPlaylists.
// Check the length with:
//     len(mockedRepo.ListPlaylistsCalls())
func (mock *RepoMock) ListPlaylistsCalls() []struct {
	Ctx    context.Context
	UserId int
} {
	var calls []struct {
		Ctx    context.Context
		UserId int
	}
	mock.lockListPlaylists.RLock()
	calls = mock.calls.ListPlaylists
	mock.lockListPlaylists.RUnlock()
	return calls
}

// RemoveSongFromPlaylist calls RemoveSongFromPlaylistFunc.
func (mock *RepoMock) RemoveSongFromPlaylist(ctx context.Context, playlistId int, songId int) error {
	if mock.RemoveSongFromPlaylistFunc == nil {
		panic("RepoMock.RemoveSongFromPlaylistFunc: method is nil but Repo.RemoveSongFromPlaylist was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		PlaylistId int
		SongId     int
	}{
		Ctx:        ctx,
		PlaylistId: playlistId,
		SongId:     songId,
	}
	mock.lockRemoveSongFromPlaylist.Lock()
	mock.calls.RemoveSongFromPlaylist = append(mock.calls.RemoveSongFromPlaylist, callInfo)
	mock.lockRemoveSongFromPlaylist.Unlock()
	return mock.RemoveSongFromPlaylistFunc(ctx, playlistId, songId)
}

// RemoveSongFromPlaylistCalls gets all the calls that were made to RemoveSongFromPlaylist.
// Check the length with:
//     len(mockedRepo.RemoveSongFromPlaylistCalls())
func (mock *RepoMock) RemoveSongFromPlaylistCalls() []struct {
	Ctx        context.Context
	PlaylistId int
	SongId     int
} {
	var calls []struct {
		Ctx        context.Context
		PlaylistId int
		SongId     int
	}
	mock.lockRemoveSongFromPlaylist.RLock()
	calls = mock.calls.RemoveSongFromPlaylist
	mock.lockRemoveSongFromPlaylist.RUnlock()
	return calls
}

// SavePlaylist calls SavePlaylistFunc.
func (mock *RepoMock) SavePlaylist(ctx context.Context, p playlistModel.AddPlaylistRequest) (playlistModel.PlaylistAggregated, error) {
	if mock.SavePlaylistFunc == nil {
		panic("RepoMock.SavePlaylistFunc: method is nil but Repo.SavePlaylist was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   playlistModel.AddPlaylistRequest
	}{
		Ctx: ctx,
		P:   p,
	}
	mock.lockSavePlaylist.Lock()
	mock.calls.SavePlaylist = append(mock.calls.SavePlaylist, callInfo)
	mock.lockSavePlaylist.Unlock()
	return mock.SavePlaylistFunc(ctx, p)
}

// SavePlaylistCalls gets all the calls that were made to SavePlaylist.
// Check the length with:
//     len(mockedRepo.SavePlaylistCalls())
func (mock *RepoMock) SavePlaylistCalls() []struct {
	Ctx context.Context
	P   playlistModel.AddPlaylistRequest
} {
	var calls []struct {
		Ctx context.Context
		P   playlistModel.AddPlaylistRequest
	}
	mock.lockSavePlaylist.RLock()
	calls = mock.calls.SavePlaylist
	mock.lockSavePlaylist.RUnlock()
	return calls
}

// UpdatePlaylist calls UpdatePlaylistFunc.
func (mock *RepoMock) UpdatePlaylist(ctx context.Context, p playlistModel.UpdatePlaylistRequest) (playlistModel.Playlist, error) {
	if mock.UpdatePlaylistFunc == nil {
		panic("RepoMock.UpdatePlaylistFunc: method is nil but Repo.UpdatePlaylist was just called")
	}
	callInfo := struct {
		Ctx context.Context
		P   playlistModel.UpdatePlaylistRequest
	}{
		Ctx: ctx,
		P:   p,
	}
	mock.lockUpdatePlaylist.Lock()
	mock.calls.UpdatePlaylist = append(mock.calls.UpdatePlaylist, callInfo)
	mock.lockUpdatePlaylist.Unlock()
	return mock.UpdatePlaylistFunc(ctx, p)
}

// UpdatePlaylistCalls gets all the calls that were made to UpdatePlaylist.
// Check the length with:
//     len(mockedRepo.UpdatePlaylistCalls())
func (mock *RepoMock) UpdatePlaylistCalls() []struct {
	Ctx context.Context
	P   playlistModel.UpdatePlaylistRequest
} {
	var calls []struct {
		Ctx context.Context
		P   playlistModel.UpdatePlaylistRequest
	}
	mock.lockUpdatePlaylist.RLock()
	calls = mock.calls.UpdatePlaylist
	mock.lockUpdatePlaylist.RUnlock()
	return calls
}