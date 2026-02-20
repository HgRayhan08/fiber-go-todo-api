package connection

import (
	"database/sql"
	"fmt"
	"log"
	"time"
	"todo-list/internal/config"

	_ "github.com/lib/pq"
)

func GetDatabaseConnection(config config.Database) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.Name,
	)

	var db *sql.DB
	var err error

	// retry maksimal 10x
	for i := 1; i <= 10; i++ {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("Attempt %d: Failed to open DB connection: %v\n", i, err)
		} else {
			err = db.Ping()
			if err == nil {
				log.Println("✅ Database Connected Successfully")
				return db
			}
			log.Printf("Attempt %d: Database not ready yet: %v\n", i, err)
		}

		log.Println("⏳ Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("❌ Could not connect to database after multiple attempts:", err)
	return nil
}
