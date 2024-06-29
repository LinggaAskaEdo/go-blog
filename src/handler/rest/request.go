package rest

type AccountLoginRequest struct {
	Username string `json:"username" example:"John"`
	Password string `json:"password" example:"$2y$10$xdTlstoTGTk5N2POd/cV6e22ByOnZmGPSjjYw9Nknd.uLO1hFuF2u"`
}
