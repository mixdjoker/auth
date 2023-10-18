package psql

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mixdjoker/auth/internal/config"
	"github.com/mixdjoker/auth/internal/model"
)

type UserStore struct {
	pool *pgxpool.Pool
}

func NewUserStore(pCfg *config.Config) *UserStore {
	dsnStr := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	dsn := fmt.Sprintf(
		dsnStr,
		os.Getenv(pCfg.Storage.Postgres.Host),
		os.Getenv(pCfg.Storage.Postgres.Port),
		os.Getenv(pCfg.User),
		os.Getenv(pCfg.Password),
		os.Getenv(pCfg.Database),
	)

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}

	return &UserStore{pool: pool}
}

func (s *UserStore) Close() {
	s.pool.Close()
}

func (s *UserStore) Create(ctx context.Context, u model.User) (int64, error) {
	var id int64
	insertBuilder := sq.Insert("public.users").
		PlaceholderFormat(sq.Dollar).
		Columns("name", "email", "password", "role").
		Values(u.Name, u.Email, u.Password, u.Role).
		Suffix("RETURNING id")

	query, args, err := insertBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	err = s.pool.QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"users_email_un\"") {
			return 0, fmt.Errorf("user with email \"%s\" already exists", u.Email)
		}

		return 0, err
	}

	return id, nil
}

func (s *UserStore) Get(ctx context.Context, id int64) (model.User, error) {
	var u model.User
	var cTime time.Time
	var uTime sql.NullTime

	selectBuilder := sq.Select("name", "email", "password", "create_at", "update_at", "role").
		From("public.users").
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		Limit(1)

	query, args, err := selectBuilder.ToSql()
	if err != nil {
		return u, err
	}

	err = s.pool.QueryRow(ctx, query, args...).Scan(&u.Name, &u.Email, &u.Password, &cTime, &uTime, &u.Role)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return u, fmt.Errorf("user with id %d not found", id)
		}

		return u, err
	}

	u.ID = id
	u.CreatedAt = cTime.Unix()
	if uTime.Valid {
		u.UpdatedAt = uTime.Time.Unix()
	}

	return u, nil
}

func (s *UserStore) Update(ctx context.Context, u model.User) error {
	updateBuilder := sq.Update("public.users").Where(sq.Eq{"id": u.ID})
	if u.Name != "" {
		updateBuilder = updateBuilder.Set("name", u.Name)
	}
	if u.Email != "" {
		updateBuilder = updateBuilder.Set("email", u.Email)
	}
	if u.Password != "" {
		updateBuilder = updateBuilder.Set("password", u.Password)
	}
	if u.Role != 0 {
		updateBuilder = updateBuilder.Set("role", u.Role)
	}
	updateBuilder = updateBuilder.Set("update_at", time.Now())

	query, args, err := updateBuilder.ToSql()
	if err != nil {
		return err
	}

	_, err = s.pool.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (s *UserStore) Delete(ctx context.Context, id int64) error {
	_, err := s.pool.Exec(ctx, "DELETE FROM public.users WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
