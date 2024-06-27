package user

import (
	"context"

	"github.com/linggaaskaedo/go-blog/src/business/entity"
)

func (u *user) GetUserByUserID(ctx context.Context, c entity.CacheControl, userID string) (entity.User, error) {
	return u.getSQLUserByID(ctx, userID)
}
