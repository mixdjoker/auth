package user_v1

import (
	"context"
	"errors"
)


func (s *serv) Delete(ctx context.Context, id int64) error {
	_, err := s.Get(ctx, id)
	if err != nil {
		return errors.New("Service.Delete: " + err.Error())
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		return s.repo.Delete(ctx, id)
	})
	if err != nil {
		return errors.New("Service.Delete: " + err.Error())
	}

	return nil
}
