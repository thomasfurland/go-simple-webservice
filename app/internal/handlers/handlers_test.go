package handlers

import (
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
