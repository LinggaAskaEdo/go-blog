package user

import (
	"context"

	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/business/entity"
	"github.com/linggaaskaedo/go-blog/src/common"
)

func (u *user) CreateUser(ctx context.Context, userEntity entity.User) (dto.UserDTO, error) {
	var userDto dto.UserDTO

	result, err := u.user.CreateUser(ctx, userEntity)
	if err != nil {
		return userDto, err
	}

	userDto.PublicID = common.MixerEncode(result.ID)
	userDto.Username = result.Username
	userDto.Email = result.Email
	userDto.Phone = result.Phone
	userDto.IsDeleted = result.IsDeleted

	return userDto, nil
}

func (u *user) GetUserByUserID(ctx context.Context, c dto.CacheControl, userID int64) (dto.UserDTO, error) {
	var userDto dto.UserDTO

	result, err := u.user.GetUserByUserID(ctx, c, userID)
	if err != nil {
		return userDto, err
	}

	userDto.PublicID = common.MixerEncode(result.ID)
	userDto.Username = result.Username
	userDto.Email = result.Email
	userDto.Phone = result.Phone
	userDto.IsDeleted = result.IsDeleted

	return userDto, nil
}
