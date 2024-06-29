package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/business/entity"
)

type DomainItf interface {
	GetUserByUserID(ctx context.Context, c dto.CacheControl, userID int64) (entity.User, error)
}

type user struct {
	logger zerolog.Logger
	redis  *redis.Client
	sql0   *sqlx.DB
	sql1   *sqlx.DB
}

type Options struct{}

func InitUserDomain(redis *redis.Client, sql0 *sqlx.DB, sql1 *sqlx.DB) DomainItf {
	return &user{
		redis: redis,
		sql0:  sql0,
		sql1:  sql1,
	}
}
