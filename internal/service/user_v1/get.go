package user_v1

import (
	"context"

	"github.com/mixdjoker/auth/internal/model"
	"github.com/pkg/errors"
)

const (
	gettingErr = "GettingError"
	notFoundErr   = "NotFoundError"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	var u *model.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		u, errTx = s.repo.Get(ctx, id)
		if errTx != nil {
			return errTx
		}
		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, gettingErr)
	}

	return u, nil
}
