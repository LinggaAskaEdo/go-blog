package division

import (
	"context"
	"database/sql"
	"time"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/business/entity"
	x "github.com/linggaaskaedo/go-blog/stdlib/error"
)

func (d *division) CreateDivision(ctx context.Context, divisionDTO dto.DivisionDTO) (entity.Division, error) {
	division := entity.Division{
		PublicID:  divisionDTO.PublicID,
		Name:      divisionDTO.Name,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}

	tx, err := d.sql0.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return division, x.Wrap("tx_create_division", err)
	}

	tx, division, err = d.createSQLDivision(tx, division)
	if err != nil {
		_ = tx.Rollback()
		return division, x.Wrap("sql_create_division", err)
	}

	if err = tx.Commit(); err != nil {
		return division, x.Wrap("comm_create_division", err)
	}

	return division, nil
}
