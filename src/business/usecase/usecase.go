package usecase

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-blog/src/business/domain"
	"github.com/linggaaskaedo/go-blog/src/business/usecase/user"
)

type Usecase struct {
	User user.UsecaseItf
}

type Options struct{}

func Init(
	logger zerolog.Logger,
	redis *redis.Client,
	sql0 *sqlx.DB,
	sql1 *sqlx.DB,
	dom *domain.Domain,
) *Usecase {
	return &Usecase{
		User: user.InitUserUsecase(
			logger,
			redis,
			sql0,
			sql1,
			dom.User,
		),
	}
}
