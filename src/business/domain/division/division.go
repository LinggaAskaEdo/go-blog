package division

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/business/entity"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type DomainItf interface {
	CreateDivision(ctx context.Context, divisionDTO dto.DivisionDTO) (entity.Division, error)
}

type division struct {
	log   zerolog.Logger
	redis *redis.Client
	sql0  *sqlx.DB
	sql1  *sqlx.DB
}

type Options struct {
}

func InitDivisionDomain(log zerolog.Logger, redis *redis.Client, sql0 *sqlx.DB, sql1 *sqlx.DB) DomainItf {
	return &division{
		log:   log,
		redis: redis,
		sql0:  sql0,
		sql1:  sql1,
	}
}
