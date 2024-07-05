package constant

const (
	// Lang Header
	LangEN string = `en`
	LangID string = `id`

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

	// Custom HTTP Header
	AppLang string = `x-app-lang`
)
