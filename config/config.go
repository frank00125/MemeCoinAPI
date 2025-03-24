package config

func init() {
	// Load environment variables
	loadEnvVars()

	// Set up database connection pool
	initDatabase()
}
