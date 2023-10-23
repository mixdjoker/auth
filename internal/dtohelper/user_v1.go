package dtohelper

import (
	"github.com/mixdjoker/auth/internal/model"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
)

func ToModelUserFromCreateRequest(req *desc.CreateRequest) *model.User {
	user := model.User{}

	if req.User.Name != nil {
		user.Name = req.User.Name.Value
	}

	if req.User.Email != nil {
		user.Email = req.User.Email.Value
	}

	if req.User.Role.Number() > 0 {
		user.Role = int(req.User.Role.Number())
	}

	return &model.User{
		Name:  req.User.Name.Value,
		Email: req.User.Email.Value,
		Role:  int(req.User.Role.Number()),
	}
}
