package storage

import (
	"context"

	"github.com/mixdjoker/auth/internal/model"
)

// UserV1Storage is an interface for user_v1 storage.
type UserV1Storage interface {
	Create(context.Context, *model.User) (int64, error)
	Get(context.Context, int64) (*model.User, error)
	Update(context.Context, *model.User) error
	Delete(context.Context, int64) error
}
