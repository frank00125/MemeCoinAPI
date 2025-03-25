package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"portto-assignment/config"
)

func main() {
	// Get database connection
	db, err := config.NewDatabaseConnectionPool()
	if err != nil {
		log.Printf("Failed to get database connection pool: %v", err)
		return
	}
	defer db.Close()

	dir, err := os.Getwd()
	if err != nil {
		log.Printf("Failed to get current working directory: %v", err)
		return
	}
	sqlFilePath := path.Join(dir, "assets", "sql", "meme_coins.sql")
	sqlBinary, err := os.ReadFile(sqlFilePath)
	if err != nil {
		log.Printf("Failed to read meme_coins.sql: %v", err)
		return
	}
	sqlStr := string(sqlBinary)

	_, err = db.ExecContext(context.Background(), sqlStr)
	if err != nil {
		log.Printf("Failed to seed meme_coins table: %v", err)
		return
	}

	fmt.Println("Successfully seeded meme_coins table!")
}
