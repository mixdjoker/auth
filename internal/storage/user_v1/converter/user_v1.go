package converter

import (
	"github.com/mixdjoker/auth/internal/model"
	"github.com/mixdjoker/auth/internal/storage/user_v1/data_model"
)

func ToModelFromStorageUserV1(user *data_model.User) *model.User {
	userModel := &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	if user.UpdatedAt.Valid {
		userModel.UpdatedAt = user.UpdatedAt.Time
	}

	return userModel
}
