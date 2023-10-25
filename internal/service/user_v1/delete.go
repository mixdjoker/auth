package user_v1

import (
	"context"
	"errors"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		u, errTx := s.repo.Get(ctx, id)
		if errTx != nil {
			return errTx
		}
		if u == nil {
			return errors.New("user not found")
		}

		return s.repo.Delete(ctx, id)
	})
	if err != nil {
		return errors.New("Service.Delete: " + err.Error())
	}

	return nil
}
