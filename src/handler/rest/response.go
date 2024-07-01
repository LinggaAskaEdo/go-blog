package rest

import (
	"github.com/linggaaskaedo/go-blog/src/business/dto"
	"github.com/linggaaskaedo/go-blog/src/business/entity"
)

type HTTPErrResp struct {
	Meta dto.Meta `json:"metadata"`
}

type HTTPEmptyResp struct {
	Meta dto.Meta `json:"metadata"`
}

type HTTPUserResp struct {
	Meta dto.Meta `json:"metadata"`
	Data UserData `json:"data"`
}

type UserData struct {
	User *entity.User `json:"user,omitempty"`
}
