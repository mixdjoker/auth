package user_v1

import (
	"context"
	"errors"

	"github.com/mixdjoker/auth/internal/model"
)

func (s *serv) Create(ctx context.Context, u *model.User) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.repo.Create(ctx, u)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.repo.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, errors.New("Service.Create: " + err.Error())
	}

	return id, nil
}
