package rest

import (
	"github.com/linggaaskaedo/go-blog/src/business/dto"
)

type HTTPErrResp struct {
	Meta dto.Meta `json:"metadata"`
}

type HTTPEmptyResp struct {
	Meta dto.Meta `json:"metadata"`
}

type HTTPDivisionResp struct {
	Meta dto.Meta     `json:"metadata"`
	Data DivisionData `json:"data"`
}

type DivisionData struct {
	Division *dto.DivisionDTO `json:"division,omitempty"`
}

type HTTPUserResp struct {
	Meta dto.Meta `json:"metadata"`
	Data UserData `json:"data"`
}

type UserData struct {
	User *dto.UserDTO `json:"user,omitempty"`
}
