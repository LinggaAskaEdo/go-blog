package constant

const (
	// Database Type
	MYSQL    string = `mysql`
	POSTGRES string = `postgres`

	// UserAgent Header
	ContentType string = `content-type`
	ContentJSON string = `application/json`

	// Cache Control Header
	CacheControl          string = `cache-control`
	CacheMustRevalidate   string = `must-revalidate`
	CacheMustDBRevalidate string = `must-db-revalidate`
)
