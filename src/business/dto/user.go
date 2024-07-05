package dto

type UserDTO struct {
	PublicID  string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	IsDeleted bool   `json:"isDeleted"`
}
