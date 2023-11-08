package dtohelper

import (
	"github.com/mixdjoker/auth/internal/model"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
)

// ToModelNewUserFromCreateRequest converts CreateRequest to model.NewUser
func ToModelNewUserFromCreateRequest(req *desc.CreateRequest) *model.NewUser {
	return &model.NewUser{
		User: &model.User{
			Name:  req.User.Name.GetValue(),
			Email: req.User.Email.GetValue(),
			Role:  int(req.User.Role.Number()),
		},
		UserCredentials: &model.UserCredentials{
			Password:        req.Password.GetValue(),
			ConfurmPassword: req.PasswordConfirm.GetValue(),
		},
	}
}

// ToModelUserFromCreateRequest converts CreateRequest to model.User
func ToModelUserFromCreateRequest(req *desc.CreateRequest) *model.User {
	return &model.User{
		Name:  req.User.Name.GetValue(),
		Email: req.User.Email.GetValue(),
		Role:  int(req.User.Role.Number()),
	}
}

// ToModelUserFromUpdateRequest converts UpdateRequest to model.User
func ToModelUserFromUpdateRequest(req *desc.UpdateRequest) *model.User {
	mUser := &model.User{
		ID: req.Id.GetValue(),
	}
	if req.User.Name != nil {
		mUser.Name = req.User.Name.GetValue()
	}
	if req.User.Email != nil {
		mUser.Email = req.User.Email.GetValue()
	}
	if req.User.Role != desc.Role_UNKNOWN {
		mUser.Role = int(req.User.Role.Number())
	}

	return mUser
}
