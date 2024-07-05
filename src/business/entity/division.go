package entity

import "database/sql"

type Division struct {
	ID        int64        `db:"id"`
	Name      string       `db:"name"`
	IsDeleted bool         `db:"is_deleted"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
