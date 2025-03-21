package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"portto-assignment/config"
)

func main() {
	// Load environment variables
	config.LoadEnvVars()

	// Inject database connection pool
	config.InitDatabase()
	connectionPool := config.GetConnection()

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v", err)
		return
	}
	memeCoinTableSQLPath := path.Join(dir, "static", "sql", "meme_coin.sql")
	memeCoinTableSQLBinary, err := os.ReadFile(memeCoinTableSQLPath)
	if err != nil {
		fmt.Printf("Failed to read meme_coin.sql: %v", err)
		return
	}
	memeCoinTableSQL := string(memeCoinTableSQLBinary)
	fmt.Println(memeCoinTableSQL)

	_, err = connectionPool.Exec(context.Background(), memeCoinTableSQL)
	if err != nil {
		fmt.Printf("Failed to seed meme_coin table: %v", err)
		return
	}

	fmt.Println("Successfully seeded meme_coin table!")
}
