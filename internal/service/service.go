package service

import (
	"context"

	"github.com/mixdjoker/auth/internal/model"
)

// UserV1Service is the interface that provides service user methods
type UserV1Service interface {
	Create(ctx context.Context, info *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, info *model.User) error
	Delete(ctx context.Context, id int64) error
}
