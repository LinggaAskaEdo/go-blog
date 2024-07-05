package user

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/linggaaskaedo/go-blog/src/business/entity"
	x "github.com/linggaaskaedo/go-blog/stdlib/errors/entity"
)

func (u *user) createSQLUser(tx *sqlx.Tx, userEntity entity.User) (*sqlx.Tx, entity.User, error) {
	row, err := tx.Exec(CreateUser, userEntity.Username, userEntity.Email, userEntity.Phone, userEntity.Division.ID, userEntity.Password, userEntity.CreatedAt)
	if err != nil {
		return tx, userEntity, x.Wrap(err, "query_create_division")
	}

	userEntity.ID, err = row.LastInsertId()
	u.log.Debug().Any("result", userEntity).Send()

	return tx, userEntity, nil
}

func (u *user) getSQLUserByID(ctx context.Context, userID int64) (entity.User, error) {
	result := entity.User{}

	sqlRow := u.sql0.QueryRowContext(ctx, GetUserByID, userID)
	if sqlRow.Err() != nil {
		return result, x.Wrap(sqlRow.Err(), "query_get_sql_user_by_id")
	}

	err := sqlRow.Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.Phone,
		&result.Division.ID,
		&result.Division.Name,
	)
	if err != nil {
		return result, x.Wrap(err, "scan_get_sql_user_by_id")
	}

	u.log.Debug().Any("result", result).Send()

	return result, nil
}
