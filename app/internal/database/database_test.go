package database

import (
	"context"
	"os"
	"testing"
)

func TestConnect_MissingDSN(t *testing.T) {
	t.Setenv("DATABASE_URL", "")
	_, err := Connect(context.Background())
	if err == nil {
		t.Fatal("expected error when DATABASE_URL is unset")
	}
}

func TestConnect_OK(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping integration test")
	}
	t.Setenv("DATABASE_URL", dsn)

	pool, err := Connect(context.Background())
	if err != nil {
		t.Fatalf("connect failed: %v", err)
	}
	defer pool.Close()
}
