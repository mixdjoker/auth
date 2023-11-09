package app

import (
	"context"
	"log"

	api_v1 "github.com/mixdjoker/auth/internal/api/user_v1"
	"github.com/mixdjoker/auth/internal/client/db"
	"github.com/mixdjoker/auth/internal/client/db/pg"
	"github.com/mixdjoker/auth/internal/client/db/transaction"
	"github.com/mixdjoker/auth/internal/closer"
	"github.com/mixdjoker/auth/internal/config"
	"github.com/mixdjoker/auth/internal/service"
	service_v1 "github.com/mixdjoker/auth/internal/service/user_v1"
	"github.com/mixdjoker/auth/internal/storage"
	storage_v1 "github.com/mixdjoker/auth/internal/storage/user_v1"
)

type serviceProvider struct {
	envConfig  config.EnvConfiger
	pgConfig   config.PgConfiger
	grpcConfig config.GrpcConfiger

	dbClient  db.Client
	txManager db.TxManager

	userStorage storage.UserV1Storage
	userService service.UserV1Service
	userImplemt *api_v1.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) EnvConfig() config.EnvConfiger {
	if s.envConfig == nil {
		envConf, err := config.NewEnvConfig()
		if err != nil {
			log.Fatalf("failed to get env config: %s", err.Error())
		}

		s.envConfig = envConf
	}

	return s.envConfig
}

func (s *serviceProvider) PGConfig() config.PgConfiger {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig(s.EnvConfig())
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GrpcConfig() config.GrpcConfiger {
	if s.grpcConfig == nil {
		cfg, err := config.NewGrpcConfig(s.EnvConfig())
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.NewClient(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %s", err.Error())
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserStorage(ctx context.Context) storage.UserV1Storage {
	if s.userStorage == nil {
		s.userStorage = storage_v1.NewRepo(s.DBClient(ctx))
	}

	return s.userStorage
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserV1Service {
	if s.userService == nil {
		s.userService = service_v1.NewService(
			s.UserStorage(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *api_v1.Implementation {
	if s.userImplemt == nil {
		s.userImplemt = api_v1.NewImplementation(s.UserService(ctx))
	}

	return s.userImplemt
}
