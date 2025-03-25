package seeds

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
)

func Seeds(db *sql.DB) {
	dir, _ := os.Getwd()
	sqlFilePath := path.Join(dir, "assets", "sql", "meme_coins.sql")
	sqlBinary, _ := os.ReadFile(sqlFilePath)
	sqlStr := string(sqlBinary)

	_, err := db.ExecContext(context.Background(), sqlStr)
	if err != nil {
		log.Printf("Failed to seed meme_coins table: %v", err)
		return
	}

	fmt.Println("Successfully seeded!")
}
