package entity

import "database/sql"

type User struct {
	ID        string `json:"id" swaggertype:"primitive,string" example:"019058b0-2953-76d4-a6cf-f89343e8c728"`
	Name      string `json:"name" swaggertype:"primitive,string" example:"Tejo"`
	Email     string `json:"email" swaggertype:"primitive,string" example:"tejo.tamvan@mss.com"`
	Phone     string `json:"phone" swaggertype:"primitive,string" example:"085712347890"`
	Division  Division
	Password  string       `json:"password" swaggertype:"primitive,string" example:"b1j2k31uob12"`
	Date      sql.NullTime `json:"date" swaggertype:"primitive,string" example:"2022-09-09T00:00:00Z"`
	CreatedAt sql.NullTime `json:"-"`
	UpdatedAt sql.NullTime `json:"-"`
	DeletedAt sql.NullTime `json:"-"`
}

type Division struct {
	ID   string `json:"id" swaggertype:"primitive,string" example:"019058b0-2953-76d4-a6cf-f89343e8c728"`
	Name string `json:"name" swaggertype:"primitive,string" example:"Admin"`
}
