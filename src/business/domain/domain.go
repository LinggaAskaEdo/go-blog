package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-blog/src/business/domain/user"
)

type Domain struct {
	User user.DomainItf
}

func Init(
	redis *redis.Client,
	sql0 *sqlx.DB,
	sql1 *sqlx.DB,
) *Domain {
	return &Domain{
		User: user.InitUserDomain(
			redis,
			sql0,
			sql1,
		),
	}
}
