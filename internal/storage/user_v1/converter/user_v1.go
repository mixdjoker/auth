package converter

import (
	"github.com/mixdjoker/auth/internal/model"
	"github.com/mixdjoker/auth/internal/storage/user_v1/data_model"
)

// ToModelUserFromRepo converts data_model.User to model.User.
func ToModelUserFromRepo(dUser *data_model.User) *model.User {
	modelUser := &model.User{
		ID:        dUser.ID,
		Name:      dUser.Name,
		Email:     dUser.Email,
		Role:      dUser.Role,
		CreatedAt: &dUser.CreatedAt,
	}

	if dUser.UpdatedAt.Valid {
		modelUser.UpdatedAt = &dUser.UpdatedAt.Time
	}

	return modelUser
}
