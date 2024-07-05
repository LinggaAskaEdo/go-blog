package entity

import "database/sql"

type User struct {
	ID        int64        `db:"id"`
	Username  string       `db:"username"`
	Email     string       `db:"email"`
	Phone     string       `db:"phone"`
	Division  Division     `db:"division"`
	Password  string       `db:"password"`
	IsDeleted bool         `db:"is_deleted"`
	CreatedAt sql.NullTime `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
