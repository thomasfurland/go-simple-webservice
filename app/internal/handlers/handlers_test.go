package handlers

import (
	"context"
	"github/thomasfurland/go-simple-webservice/internal/database"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	w := httptest.NewRecorder()

	homeHandler(w, req)

	res := w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.StatusCode)
	}

	body := w.Body.String()

	text := "hello world!"

	if body != text {
		t.Fatalf("expected '%q', got %q", text, body)
	}
}

func TestDatabaseHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/db", nil)

	pool, err := database.Connect(context.Background())
	if err != nil {
		t.Fatalf("failed to start database")
	}

	h := &Handler{pool: pool}

	w := httptest.NewRecorder()

	h.dbHandler(w, req)

	res := w.Result()

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", res.StatusCode)
	}

	body := w.Body.String()

	text := "ok"

	if body != text {
		t.Fatalf("expected '%q', got %q", text, body)
	}
}
