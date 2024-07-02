package division

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"

	div "github.com/linggaaskaedo/go-blog/src/business/domain/division"
	"github.com/linggaaskaedo/go-blog/src/business/dto"
)

type UsecaseItf interface {
	CreateDivision(ctx context.Context, division dto.DivisionDTO) (dto.DivisionDTO, error)
}

type division struct {
	log      zerolog.Logger
	redis    *redis.Client
	sql0     *sqlx.DB
	sql1     *sqlx.DB
	division div.DomainItf
}

type Options struct {
}

func InitDivisionUsecase(log zerolog.Logger, redis *redis.Client, sql0 *sqlx.DB, sql1 *sqlx.DB, d div.DomainItf) UsecaseItf {
	return &division{
		log:      log,
		redis:    redis,
		sql0:     sql0,
		sql1:     sql1,
		division: d,
	}
}
