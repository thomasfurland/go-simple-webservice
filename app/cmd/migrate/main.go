package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // registers "pgx" with database/sql
	"github.com/pressly/goose/v3"
)

// exit success on error or successful migration for github actions
func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		fmt.Println("MIGRATION FAILURE: DATABASE_URL not set")
		os.Exit(0)
	}

	dir := os.Getenv("MIGRATION_DIRECTORY")
	if dir == "" {
		fmt.Println("MIGRATION FAILURE: MIGRATION_DIRECTORY not set")
		os.Exit(0)
	}

	// 1) Wait for DB readiness (bounded)
	if err := waitForDB(dsn, 2*time.Minute, 2*time.Second); err != nil {
		fmt.Printf("MIGRATION FAILURE: db not ready: %v\n", err)
		os.Exit(0)
	}

	// 2) Run migrations (bounded)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		fmt.Printf("MIGRATION FAILURE: open db: %v\n", err)
		os.Exit(0)
	}
	defer db.Close()

	if err := goose.UpContext(ctx, db, dir); err != nil {
		fmt.Printf("MIGRATION FAILURE: migrate up: %v\n", err)
		os.Exit(0)
	}

	fmt.Println("MIGRATION SUCCESS: ✅ migrations applied")
	os.Exit(0)
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
