package user_v1

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/mixdjoker/auth/internal/client/db"
	"github.com/mixdjoker/auth/internal/model"
	"github.com/mixdjoker/auth/internal/storage"
	"github.com/mixdjoker/auth/internal/storage/user_v1/converter"
	"github.com/mixdjoker/auth/internal/storage/user_v1/data_model"
)

const (
	userTable = "public.users"

	idColumn         = "user_id"
	nameColumn       = "name"
	emailColumn      = "email"
	passwordColumn   = "password"
	roleColumn       = "role_id"
	ctraetedAtColumn = "created_at"
	updatedAtColumn  = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepo(db db.Client) storage.UserV1Storage {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, u *model.User) (int64, error) {
	retStr := fmt.Sprintf("RETURNING %s", idColumn)
	insertBuilder := sq.Insert(userTable).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(u.Name, u.Email, u.Password, u.Role).
		Suffix(retStr)

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_v1.Create",
		QueryRaw: query,
	}

	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	selectBuilder := sq.Select(
		nameColumn,
		emailColumn,
		ctraetedAtColumn,
		updatedAtColumn,
		roleColumn).
		From(userTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_v1.Get",
		QueryRaw: query,
	}

	var dUser data_model.User
	err = r.db.DB().ScanOneContext(ctx, &dUser, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToModelFromStorageUserV1(&dUser), nil
}

func (r *repo) Update(ctx context.Context, u *model.User) error {
	curUser, err := r.Get(ctx, u.ID)
	if err != nil {
		return err
	}

	updateBuilder := sq.Update(userTable).
		Where(sq.Eq{idColumn: u.ID}).
		PlaceholderFormat(sq.Dollar)
	if u.Name != curUser.Name {
		updateBuilder = updateBuilder.Set(nameColumn, u.Name)
	}
	if u.Email != curUser.Email {
		updateBuilder = updateBuilder.Set(emailColumn, u.Email)
	}
	if u.Password != curUser.Password {
		updateBuilder = updateBuilder.Set(passwordColumn, u.Password)
	}
	if u.Role != curUser.Role {
		updateBuilder = updateBuilder.Set(roleColumn, u.Role)
	}
	updateBuilder = updateBuilder.Set(updatedAtColumn, time.Now())
	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "user_v1.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	deleteBuilder := sq.Delete(userTable).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := deleteBuilder.ToSql()

	q := db.Query{
		Name:     "user_v1.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
