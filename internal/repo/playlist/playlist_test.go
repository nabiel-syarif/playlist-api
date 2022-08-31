package playlist

import (
	"reflect"
	"testing"

	db "github.com/nabiel-syarif/playlist-api/pkg/db"
	"github.com/pashagolub/pgxmock"
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
			desc: "case 1 -> success when init playlist repo",
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
