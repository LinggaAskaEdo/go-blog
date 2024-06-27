package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-blog/src/business/domain/user"
)

type Domain struct {
	User user.DomainItf
}

func Init(
	logger zerolog.Logger,
	redis *redis.Client,
	sql0 *sqlx.DB,
	sql1 *sqlx.DB,
) *Domain {
	return &Domain{
		User: user.InitUserDomain(
			logger,
			redis,
			sql0,
			sql1,
		),
	}
}
