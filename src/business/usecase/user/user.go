package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	usr "github.com/linggaaskaedo/go-blog/src/business/domain/user"
	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/business/entity"
)

type UsecaseItf interface {
	GetUserByUserID(ctx context.Context, c dto.CacheControl, userID int64) (entity.User, error)
}

type user struct {
	redis *redis.Client
	sql0  *sqlx.DB
	sql1  *sqlx.DB
	user  usr.DomainItf
}

type Options struct{}

func InitUserUsecase(redis *redis.Client, sql0 *sqlx.DB, sql1 *sqlx.DB, u usr.DomainItf) UsecaseItf {
	return &user{
		redis: redis,
		sql0:  sql0,
		sql1:  sql1,
		user:  u,
	}
}
