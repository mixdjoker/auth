package user_v1

import (
	"context"

	"github.com/pkg/errors"
)

const (
	deleteErr = "DeleteError"
)

func (s *serv) Delete(ctx context.Context, id int64) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		_, errTx := s.repo.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return s.repo.Delete(ctx, id)
	})
	if err != nil {
		return errors.Wrap(err, deleteErr)
	}

	return nil
}
