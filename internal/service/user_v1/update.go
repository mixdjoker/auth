package user_v1

import (
	"context"
	"errors"

	"github.com/mixdjoker/auth/internal/model"
)

func (s *serv) Update(ctx context.Context, info *model.User) error {
	exist, err := s.Get(ctx, info.ID)
	if err != nil {
		return errors.New("Service.Update: " + err.Error())
	}

	if exist == nil {
		return errors.New("Service.Update: user not found")
	}

	updateExistInfo(exist, info)

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
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
