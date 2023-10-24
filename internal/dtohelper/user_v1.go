package dtohelper

import (
	"github.com/mixdjoker/auth/internal/model"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
)

func ToModelUserFromCreateRequest(req *desc.CreateRequest) *model.User {
	return &model.User{
		Name:     req.User.Name.Value,
		Email:    req.User.Email.Value,
		Password: req.Password.Value,
		Role:     int(req.User.Role.Number()),
	}
}

func ToModelUserFromUpdateRequest(req *desc.UpdateRequest) *model.User {
	var (
		name  string
		email string
		role  int
	)
	if req.User.Name != nil {
		name = req.User.Name.Value
	}

	if req.User.Email != nil {
		email = req.User.Email.Value
	}

	if req.User.Role != desc.Role_UNKNOWN {
		role = int(req.User.Role.Number())
	}

	return &model.User{
		ID:    req.Id.Value,
		Name:  name,
		Email: email,
		Role:  role,
	}
}
