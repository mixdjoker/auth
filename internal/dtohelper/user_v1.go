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
