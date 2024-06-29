package user

import (
	"context"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/business/entity"
)

func (u *user) GetUserByUserID(ctx context.Context, c dto.CacheControl, userID int64) (entity.User, error) {
	return u.getSQLUserByID(ctx, userID)
}
