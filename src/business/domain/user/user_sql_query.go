package user

const (
	GetUserByID = `SELECT
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
		user.id = ?`
)