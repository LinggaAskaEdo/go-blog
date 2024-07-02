package entity

import "database/sql"

type Division struct {
	ID        int64        `json:"id"`
	PublicID  string       `json:"public_id"`
	Name      string       `json:"name"`
	CreatedAt sql.NullTime `json:"created_at"`
	UpdatedAt sql.NullTime `json:"-"`
}
