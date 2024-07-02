package rest

type AccountLoginRequest struct {
	Data CreateUserData `json:"data"`
}

type CreateUserData struct {
	Username string `json:"username" example:"John"`
	Password string `json:"password" example:"$2y$10$xdTlstoTGTk5N2POd/cV6e22ByOnZmGPSjjYw9Nknd.uLO1hFuF2u"`
}

type DivisionCreateRequest struct {
	Data CreateDivisionData `json:"data"`
}

type CreateDivisionData struct {
	Division *DivisionDataPayload `json:"division,omitempty"`
}

type DivisionDataPayload struct {
	Name string `json:"name"`
}
