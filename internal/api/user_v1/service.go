package user_v1

import (
	"github.com/mixdjoker/auth/internal/service"
	desc "github.com/mixdjoker/auth/pkg/user_v1"
)

// Implementation of UserV1Server interface
type Implementation struct {
	desc.UnimplementedUser_V1Server
	userService service.UserV1Service
}

// NewImplementation returns a new instance of UserV1Server implementation
func NewImplementation(userV1Service service.UserV1Service) *Implementation {
	return &Implementation{
		userService: userV1Service,
	}
}
