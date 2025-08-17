package main

import (
	"log"
	"social/internal/db"
	"social/internal/env"
	"social/internal/storage"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file, using environment variables")
	}
}

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8081"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "host=db port=5432 user=admin password=adminpassword dbname=social sslmode=disable"),
			maxOpenConns: env.GetEnvInt("DB_MAX_OPEN_CONNS", 10),
			maxIdleConns: env.GetEnvInt("DB_MAX_IDLE_CONNS", 5),
			maxIdleTime:  time.Duration(env.GetEnvInt("DB_MAX_IDLE_TIME", 60)) * time.Second,
		},
	}

	log.Printf("Starting application with config: addr=%s", cfg.addr)

	// Wait for database to be ready
	log.Println("Waiting for database to be ready...")
	time.Sleep(5 * time.Second)

	// Connect to database
	database, err := db.NewDB(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.Close()

	log.Println("Database connection established")

	// Initialize storage
	store := storage.NewStorage(database)

	// Create application
	app := &application{
		config:  cfg,
		dbLayer: store,
	}

	// Setup router
	router := app.mount()

	log.Printf("Server starting on %s", cfg.addr)
	log.Fatal(app.run(router))
}
