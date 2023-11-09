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

// NewRepo returns a new instance of storage.UserV1Storage
func NewRepo(db db.Client) storage.UserV1Storage {
	return &repo{db: db}
}

// Create creates a new user in the database
func (r *repo) Create(ctx context.Context, u *model.NewUser) (int64, error) {
	retStr := fmt.Sprintf("RETURNING %s", idColumn)
	insertBuilder := sq.Insert(userTable).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(u.User.Name, u.User.Email, u.UserCredentials.Password, u.User.Role).
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

// Get returns a user from the database
func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	selectBuilder := sq.Select(
		idColumn,
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

	return converter.ToModelUserFromRepo(&dUser), nil
}

// Update updates a user in the database
func (r *repo) Update(ctx context.Context, u *model.User) error {
	updateBuilder := sq.Update(userTable).
		Where(sq.Eq{idColumn: u.ID}).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, u.Name).
		Set(emailColumn, u.Email).
		Set(roleColumn, u.Role).
		Set(updatedAtColumn, time.Now())

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

// Delete deletes a user from the database
func (r *repo) Delete(ctx context.Context, id int64) error {
	deleteBuilder := sq.Delete(userTable).
		Where(sq.Eq{idColumn: id}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := deleteBuilder.ToSql()
	if err != nil {
		return err
	}

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

// GetUserByEmail returns a user from the database by email
func (r *repo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	qBuilder := sq.Select(
		idColumn,
		nameColumn,
		emailColumn,
		ctraetedAtColumn,
		updatedAtColumn,
		roleColumn).
		From(userTable).
		Where(sq.Eq{email: email}).
		PlaceholderFormat(sq.Dollar)

	query, args, err := qBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_v1.GetUserByEmail",
		QueryRaw: query,
	}

	var dUser data_model.User
	err = r.db.DB().ScanOneContext(ctx, &dUser, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToModelUserFromRepo(&dUser), nil
}
