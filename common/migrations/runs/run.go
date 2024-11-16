package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func main() {
	dsn := "host=localhost user=postgres password=1234 dbname=app_rice_wine port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	folderIncludeFileMigrations := "common/migrations"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	files, err := os.ReadDir(folderIncludeFileMigrations)
	if err != nil {
		log.Fatal("Error reading migrations folder:", err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".sql" {
			filePath := filepath.Join(folderIncludeFileMigrations, file.Name())

			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Error reading file %s: %v", file.Name(), err)
				continue
			}

			_, err = db.Exec(string(content))
			if err != nil {
				log.Printf("Error executing migration %s: %v", file.Name(), err)
				continue
			}

			fmt.Printf("Successfully executed migration: %s\n", file.Name())
		}
	}

	fmt.Println("Migration completed")
}
