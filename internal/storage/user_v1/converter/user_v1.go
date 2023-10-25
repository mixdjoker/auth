package converter

import (
	"github.com/mixdjoker/auth/internal/model"
	"github.com/mixdjoker/auth/internal/storage/user_v1/data_model"
)

func ToModelUserFromRepo(dUser *data_model.User) *model.User {
	userModel := &model.User{
		ID:        dUser.ID,
		Name:      dUser.Name,
		Email:     dUser.Email,
		Password:  dUser.Password,
		Role:      dUser.Role,
		CreatedAt: dUser.CreatedAt,
	}

	if dUser.UpdatedAt.Valid {
		userModel.UpdatedAt = &dUser.UpdatedAt.Time
	}

	return userModel
}
