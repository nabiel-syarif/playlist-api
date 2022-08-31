package song

import (
	"context"
	"reflect"
	"testing"
	"time"

	modelSong "github.com/nabiel-syarif/playlist-api/internal/model/song"
	db "github.com/nabiel-syarif/playlist-api/pkg/db"
	pgxmock "github.com/pashagolub/pgxmock"
)

func TestNew(t *testing.T) {

	dbConn, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("%v\n", err.Error())
	}
	type args struct {
		db db.DB
	}
	testCases := []struct {
		desc string
		args args
		want Repo
	}{
		{
			desc: "case 1 -> success when init song repo",
			args: args{
				db: dbConn,
			},
			want: &repo{
				db: dbConn,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			if actual := New(dbConn); !reflect.DeepEqual(actual, tC.want) {
				t.Fatalf("Want : %v but got %v", tC.want, actual)
			}
		})
	}
}

func TestRepo_InsertSong(t *testing.T) {
	type args struct {
		ctx  context.Context
		song modelSong.InsertSongRequest
	}

	ctx := context.Background()
	cols := []string{"song_id", "title", "performer", "genre", "duration", "created_at", "updated_at"}

	testCases := []struct {
		desc          string
		args          args
		getRepo       func(modelSong.InsertSongRequest) Repo
		wantErr       bool
		error         interface{}
		equalityCheck func(modelSong.Song, args)
	}{
		{
			desc: "case 1 -> success insert a song",
			args: args{
				ctx: ctx,
				song: modelSong.InsertSongRequest{
					Title:     "tracing that dream",
					Performer: "Yoasobi",
					Genre:     "pop",
					Duration:  120,
				},
			},
			getRepo: func(song modelSong.InsertSongRequest) Repo {
				dbConn, err := pgxmock.NewConn()
				if err != nil {
					t.Fatalf("Err init db conn: %v\n", err)
				}
				rows := pgxmock.NewRows([]string{"song_id"})
				songId := 1
				rows.AddRow(songId)
				dbConn.ExpectQuery(`INSERT INTO songs\(title, performer, genre, duration\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING song_id`).WithArgs(song.Title, song.Performer, song.Genre, song.Duration).WillReturnRows(rows)

				rows = pgxmock.NewRows(cols)
				rows.AddRow(songId, song.Title, song.Performer, song.Genre, song.Duration, time.Now(), time.Now())
				dbConn.ExpectQuery(`SELECT \* FROM songs WHERE song_id = \$1 LIMIT 1`).WithArgs(songId).WillReturnRows(rows)

				return New(dbConn)
			},
			wantErr: false,
			equalityCheck: func(s modelSong.Song, args args) {
				if s.Title != args.song.Title || s.Performer != args.song.Performer || s.Genre != args.song.Genre || s.Duration != args.song.Duration {
					t.Fatalf("Want %v but got %v", args.song, s)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			song, err := tC.getRepo(tC.args.song).InsertSong(ctx, tC.args.song)
			if err != nil {
				if !tC.wantErr {
					t.Fatalf("Repo.InsertSong() error : %v, should not return error\n", err)
				} else if !reflect.DeepEqual(tC.error, err) {
					t.Fatalf("Repo.InsertSong() actual error : %v, want error : %v", err, tC.error)
				}
			}
			tC.equalityCheck(song, tC.args)
		})
	}
}
