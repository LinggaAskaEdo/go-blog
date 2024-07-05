package user

const (
	CreateUser = `
		INSERT INTO user 
		(
			username, 
			email,
			phone,
			division_id,
			password,
			created_at
		) 
		VALUES 
			(?,?,?,?,?,?);
	`

	GetUserByID = `
		SELECT
			user.id,
			user.name,
			user.email,
			user.phone,
			division.id,
			division.name,
			user.password,
			user.date,
			user.createdAt,
			user.updatedAt,
			user.deletedAt
		FROM user
		JOIN division
			ON user.division_id = division.id
		WHERE
			user.id = ?
	`
)
