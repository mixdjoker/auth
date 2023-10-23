package user_v1

import (
	"context"
	"errors"

	"github.com/mixdjoker/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, u *model.User) (int64, error) {
	return 0, errors.New("not implemented")
}
