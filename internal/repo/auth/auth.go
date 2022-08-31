package auth

import (
	"context"

	"github.com/jackc/pgx/v4"
	authModel "github.com/nabiel-syarif/playlist-api/internal/model/auth"
	"github.com/nabiel-syarif/playlist-api/pkg/db"
)

type repo struct {
	db db.DB
}

type Repo interface {
	GetUserByEmail(ctx context.Context, email string) (authModel.User, error)
	GetUsersByUsersId(ctx context.Context, userIds []int) ([]authModel.User, error)
	GetUserByUserId(ctx context.Context, userId int) (authModel.User, error)
	Register(ctx context.Context, newUser authModel.UserRegistration) (authModel.User, error)
}

var _ Repo = &repo{}

func New(db db.DB) Repo {
	return &repo{
		db: db,
	}
}

func (repo *repo) GetUserByEmail(ctx context.Context, email string) (authModel.User, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM users WHERE email = $1", email)

	var user authModel.User

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo *repo) Register(ctx context.Context, newUser authModel.UserRegistration) (authModel.User, error) {
	res := repo.db.QueryRow(ctx, "INSERT INTO users(name, email, password) VALUES ($1, $2, $3) RETURNING user_id", newUser.Name, newUser.Email, newUser.Password)

	var user authModel.User

	var id int
	err := res.Scan(&id)
	if err != nil {
		return user, err
	}

	row := repo.db.QueryRow(context.Background(), "SELECT * FROM users WHERE user_id = $1 LIMIT 1", id)
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repo) GetUsersByUsersId(ctx context.Context, userId []int) ([]authModel.User, error) {
	if len(userId) == 0 {
		return nil, pgx.ErrNoRows
	}
	params := make([]interface{}, 0)
	sql := "SELECT * FROM users WHERE user_id IN ("
	for i, v := range userId {
		sql += `$` + string(i+1)
		params = append(params, v)
	}
	sql += ")"

	rows, err := r.db.Query(ctx, sql, params...)
	if err != nil {
		return nil, err
	}

	users := make([]authModel.User, 0)
	for rows.Next() {
		var user authModel.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *repo) GetUserByUserId(ctx context.Context, userId int) (authModel.User, error) {
	sql := "SELECT * FROM users WHERE user_id = $1"
	row := r.db.QueryRow(ctx, sql, userId)
	var user authModel.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return authModel.User{}, err
	}
	return user, nil
}
