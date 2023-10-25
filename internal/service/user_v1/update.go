package user_v1

import (
	"context"
	"errors"

	"github.com/mixdjoker/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.User) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		exist, errTx := s.repo.Get(ctx, info.ID)
		if errTx != nil {
			return errTx
		}

		if exist == nil {
			return errors.New("user not found")
		}

		updateExistInfo(exist, info)

		errTx = s.repo.Update(ctx, exist)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return errors.New("Service.Update: " + err.Error())
	}

	return nil
}

func updateExistInfo(exist *model.User, info *model.User) {
	if info.Name != "" {
		exist.Name = info.Name
	}

	if info.Email != "" {
		exist.Email = info.Email
	}

	if info.Role != 0 {
		exist.Role = info.Role
	}
}
