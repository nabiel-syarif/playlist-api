package auth

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"
	modelAuth "github.com/nabiel-syarif/playlist-api/internal/model/auth"
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
			desc: "case 1 -> success when init auth repo",
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

func TestRepo_GetUserByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}

	ctx := context.Background()
	email := "blabla@gmail.com"
	cols := []string{"user_id", "name", "email", "password", "created_at", "updated_at"}

	testCases := []struct {
		desc          string
		args          args
		getRepo       func() Repo
		wantErr       bool
		error         interface{}
		equalityCheck func(modelAuth.User)
	}{
		{
			desc: "case 1 -> success found a user associated with an email",
			args: args{
				ctx:   ctx,
				email: email,
			},
			getRepo: func() Repo {
				dbConn, err := pgxmock.NewConn()
				if err != nil {
					t.Fatalf("Err init db conn: %v\n", err)
				}
				rows := pgxmock.NewRows(cols)
				rows.AddRow(1, "nabiel", email, "password", time.Now(), time.Now())
				dbConn.ExpectQuery("SELECT \\* FROM users WHERE email = \\$1").WithArgs(email).WillReturnRows(rows)
				return New(dbConn)
			},
			wantErr: false,
			equalityCheck: func(u modelAuth.User) {
				if u.Email != email {
					t.Fatalf("Want user with email %v, but got %v", email, u.Email)
				}
			},
		},
		{
			desc: "case 2 -> should return error when user not found",
			args: args{
				ctx:   ctx,
				email: email,
			},
			getRepo: func() Repo {
				dbConn, err := pgxmock.NewConn()
				if err != nil {
					t.Fatalf("Err init db conn: %v\n", err)
				}
				dbConn.ExpectQuery("SELECT \\* FROM users WHERE email = \\$1").WithArgs(email).WillReturnError(pgx.ErrNoRows)
				return New(dbConn)
			},
			wantErr: true,
			error:   pgx.ErrNoRows,
			equalityCheck: func(u modelAuth.User) {
				if u.Email != "" {
					t.Fatalf("Want user with email %v, but got %v", email, u.Email)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			user, err := tC.getRepo().GetUserByEmail(ctx, email)
			if err != nil {
				if !tC.wantErr {
					t.Fatalf("Repo.GetUserByEmail() error : %v, should not return error\n", err)
				} else if !reflect.DeepEqual(tC.error, err) {
					t.Fatalf("Repo.GetUserByEmail() actual error : %v, want error : %v", err, tC.error)
				}
			}
			tC.equalityCheck(user)
		})
	}
}

func TestRepo_Register(t *testing.T) {
	type args struct {
		ctx  context.Context
		user modelAuth.UserRegistration
	}

	ctx := context.Background()
	user := modelAuth.UserRegistration{
		Name:     "nabiel",
		Email:    "nabiel@gmail.com",
		Password: "blabla",
	}
	cols := []string{"user_id", "name", "email", "password", "created_at", "updated_at"}

	testCases := []struct {
		desc          string
		args          args
		getRepo       func() Repo
		wantErr       bool
		error         interface{}
		equalityCheck func(modelAuth.User)
	}{
		{
			desc: "case 1 -> success found a user associated with an email",
			args: args{
				ctx:  ctx,
				user: user,
			},
			getRepo: func() Repo {
				dbConn, err := pgxmock.NewConn()
				if err != nil {
					t.Fatalf("Err init db conn: %v\n", err)
				}
				rows := pgxmock.NewRows([]string{"user_id"})
				userId := 1
				rows.AddRow(userId)
				dbConn.ExpectQuery("INSERT INTO users\\(name, email, password\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING user_id").WithArgs(user.Name, user.Email, user.Password).WillReturnRows(rows)
				rows = pgxmock.NewRows(cols)
				rows.AddRow(userId, user.Name, user.Email, user.Password, time.Now(), time.Now())
				dbConn.ExpectQuery(`SELECT \* FROM users WHERE user_id = \$1 LIMIT 1`).WithArgs(userId).WillReturnRows(rows)

				return New(dbConn)
			},
			wantErr: false,
			equalityCheck: func(u modelAuth.User) {
				if u.Name != user.Name || u.Email != user.Email || u.Password != user.Password {
					t.Fatalf("Want %v but got %v", user, u)
				}
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			user, err := tC.getRepo().Register(ctx, user)
			if err != nil {
				if !tC.wantErr {
					t.Fatalf("Repo.Register() error : %v, should not return error\n", err)
				} else if !reflect.DeepEqual(tC.error, err) {
					t.Fatalf("Repo.Register() actual error : %v, want error : %v", err, tC.error)
				}
			}
			tC.equalityCheck(user)
		})
	}
}
