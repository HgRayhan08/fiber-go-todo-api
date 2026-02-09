package connection

import (
	"database/sql"
	"fmt"
	"log"
	"todo-list/internal/config"

	_ "github.com/lib/pq"
)

func GetDatabaseConnection(config config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to Open Connection Database:", err)
	}
	if db == nil {
		log.Fatal("failed to create goqu database instance")
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to Ping Database:", err)
	}
	log.Println("😑 Database Connected Successfully")
	return db
}
