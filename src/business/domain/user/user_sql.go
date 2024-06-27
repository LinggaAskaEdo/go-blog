package user

import (
	"context"

	"github.com/linggaaskaedo/go-blog/src/business/entity"
)

func (u *user) getSQLUserByID(ctx context.Context, userID string) (entity.User, error) {
	result := entity.User{}

	sqlRow := u.sql0.QueryRowContext(ctx, GetUserByID, userID)
	if sqlRow.Err() != nil {
		return result, sqlRow.Err()
	}

	err := sqlRow.Scan(
		&result.ID,
		&result.Name,
		&result.Email,
		&result.Phone,
		&result.Division.ID,
		&result.Division.Name,
	)
	if err != nil {
		return result, err
	}

	return result, nil
}
