package division

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-blog/src/business/entity"
	x "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"
)

func (d *division) CreateDivision(ctx context.Context, divisionEntity entity.Division) (entity.Division, error) {
	tx, err := d.sql0.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return divisionEntity, x.Wrap(err, "tx_create_division")
	}

	tx, divisionEntity, err = d.createSQLDivision(tx, divisionEntity)
	if err != nil {
		_ = tx.Rollback()
		return divisionEntity, x.Wrap(err, "sql_create_division")
	}

	if err = tx.Commit(); err != nil {
		return divisionEntity, x.Wrap(err, "commit_create_division")
	}

	return divisionEntity, nil
}

func (d *division) GetDivisioByID(ctx context.Context, divisionID int64) (entity.Division, error) {
	result, err := d.getSQLDivisionByID(ctx, divisionID)
	if err != nil {
		return result, err
	}

	return result, nil
}
