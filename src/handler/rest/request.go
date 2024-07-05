package rest

type DivisionCreateRequest struct {
	Data CreateDivisionData `json:"data"`
}

type CreateDivisionData struct {
	Division *DivisionDataPayload `json:"division,omitempty"`
}

type DivisionDataPayload struct {
	Name string `json:"name"`
}

type UserCreateRequest struct {
	Data CreateUserData `json:"data"`
}

type CreateUserData struct {
	User *UserDataPayload `json:"user,omitempty"`
}

type UserDataPayload struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	DivisionID string `json:"divisionID"`
	Password   string `json:"password"`
}
