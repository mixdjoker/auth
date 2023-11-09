package user_v1

import (
	"context"

	"github.com/mixdjoker/auth/internal/model"
	"github.com/pkg/errors"
)

const (
	// CreateErr is the error when create user.
	createErr    = "CreationError"
	validatorErr = "ValidationError"
)

func (s *serv) Create(ctx context.Context, u *model.NewUser) (int64, error) {
	validateErr := u.Validate(localValidator)
	if validateErr != nil {
		return 0, errors.Wrap(validateErr, validatorErr)
	}

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
		return 0, errors.Wrap(err, createErr)
	}

	return id, nil
}

func localValidator(u *model.NewUser) error {
	if u.User.Name == "" {
		return errors.New("Name is required")
	}

	if u.User.Email == "" {
		return errors.New("Email is required")
	}

	if u.User.Role == 0 {
		return errors.New("Role is required")
	}

	if u.UserCredentials.Password == "" {
		return errors.New("Password is required")
	}

	if u.UserCredentials.ConfurmPassword == "" {
		return errors.New("Password confirm is required")
	}

	if u.UserCredentials.Password != u.UserCredentials.ConfurmPassword {
		return errors.New("Password and Password confirm must be the same")
	}

	return nil
}
