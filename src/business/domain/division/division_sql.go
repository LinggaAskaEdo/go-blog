package division

import (
	"github.com/jmoiron/sqlx"

	"github.com/linggaaskaedo/go-blog/src/business/entity"
	x "github.com/linggaaskaedo/go-blog/stdlib/error"
)

func (d *division) createSQLDivision(tx *sqlx.Tx, division entity.Division) (*sqlx.Tx, entity.Division, error) {
	rows, err := tx.Exec(CreateDivision, division.PublicID, division.Name, division.CreatedAt)
	if err != nil {
		return tx, division, x.Wrap("TxCreateDivision", err)
	}

	division.ID, err = rows.LastInsertId()
	d.log.Debug().Any("result", division).Send()

	return tx, division, nil
}
