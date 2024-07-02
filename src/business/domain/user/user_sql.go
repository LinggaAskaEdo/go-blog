package user

import (
	"context"

	"github.com/linggaaskaedo/go-blog/src/business/entity"
	x "github.com/linggaaskaedo/go-blog/stdlib/error"
)

func (u *user) getSQLUserByID(ctx context.Context, userID int64) (entity.User, error) {
	result := entity.User{}

	sqlRow := u.sql0.QueryRowContext(ctx, GetUserByID, userID)
	if sqlRow.Err() != nil {
		return result, x.Wrap("getSQLUserByID", sqlRow.Err())
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
		return result, x.Wrap("getSQLUserByID", err)
	}

	return result, nil
}
