package division

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/linggaaskaedo/go-blog/src/business/entity"
	commonerr "github.com/linggaaskaedo/go-blog/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"
)

func (d *division) createSQLDivision(tx *sqlx.Tx, divisionEntity entity.Division) (*sqlx.Tx, entity.Division, error) {
	row, err := tx.Exec(CreateDivision, divisionEntity.Name, divisionEntity.CreatedAt)
	if err != nil {
		return tx, divisionEntity, x.Wrap(err, "query_create_division")
	}

	divisionEntity.ID, err = row.LastInsertId()
	d.log.Debug().Any("result", divisionEntity).Send()

	return tx, divisionEntity, nil
}

func (d *division) getSQLDivisionByID(ctx context.Context, divisionID int64) (entity.Division, error) {
	var result entity.Division

	// err := d.sql0.GetContext(ctx, &result, GetDivisionByID, divisionID)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		return result, x.WrapWithCode("query_get_division_by_id", err, http.StatusNotFound)
	// 	}

	// 	return result, x.Wrap("query_get_division_by_id", err)
	// }

	row := d.sql0.QueryRowContext(ctx, GetDivisionByID, divisionID)

	if err := row.Scan(
		&result.ID,
		&result.Name,
		&result.IsDeleted,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return result, x.WrapWithCode(err, commonerr.CodeHTTPNotFound, "query_get_division_by_id")
		}

		return result, x.Wrap(err, "scan_get_division_by_id")
	}

	d.log.Debug().Any("result", result).Send()

	return result, nil
}
