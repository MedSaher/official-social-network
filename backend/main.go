package main

import (
	"log"
	"social_network/internal/infrastructure/db/sqlite"
)

	func main() {
		dbPath := "./social_network/social_network.db"
		migrationsPath := "./internal/infrastructure/db/migrations"
	
		db, err := sqlite.ConnectDB(dbPath)
		if err != nil {
			log.Fatal("DB connection error:", err)
		}
	
		if err := sqlite.RunMigrations(db, migrationsPath); err != nil {
			log.Fatal("Migration error:", err)
		}
	
		// db ready to use...
	}
	