package division

const (
	CreateDivision = `
		INSERT INTO division 
		(
			name, 
			created_at
		) 
		VALUES 
			(?,?);
	`

	GetDivisionByID = `
		SELECT 
			id, 
			name, 
			is_deleted, 
			created_at, 
			updated_at, 
			deleted_at 
		FROM 
			division 
		WHERE 
			id = ?
	`
)
