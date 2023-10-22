package user_v1

import (
	"github.com/mixdjoker/auth/internal/client/db"
	"github.com/mixdjoker/auth/internal/service"
	"github.com/mixdjoker/auth/internal/storage"
)

type serv struct {
	repo      storage.UserV1Storage
	txManager db.TxManager
}

func NewService(repo storage.UserV1Storage, txManager db.TxManager) service.UserV1Service {
	return &serv{
		repo:      repo,
		txManager: txManager,
	}
}
