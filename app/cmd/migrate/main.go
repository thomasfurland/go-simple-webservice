package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // registers "pgx" with database/sql
	"github.com/pressly/goose/v3"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	dir := os.Getenv("MIGRATION_DIRECTORY")
	if dir == "" {
		log.Fatal("MIGRATION_DIRECTORY not set")
	}

	// 1) Wait for DB readiness (bounded)
	if err := waitForDB(dsn, 2*time.Minute, 2*time.Second); err != nil {
		log.Fatalf("db not ready: %v", err)
	}

	// 2) Run migrations (bounded)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	if err := goose.UpContext(ctx, db, dir); err != nil {
		log.Fatalf("migrate up: %v", err)
	}

	log.Println("✅ migrations applied")
}

func waitForDB(dsn string, maxWait, step time.Duration) error {
	deadline := time.Now().Add(maxWait)
	for {
		ctx, cancel := context.WithTimeout(context.Background(), step)
		ok := pingOnce(ctx, dsn)
		cancel()
		if ok {
			return nil
		}
		if time.Now().After(deadline) {
			return errors.New("timeout waiting for database")
		}
		time.Sleep(step)
	}
}

func pingOnce(ctx context.Context, dsn string) bool {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return false
	}
	defer db.Close()
	return db.PingContext(ctx) == nil
}
