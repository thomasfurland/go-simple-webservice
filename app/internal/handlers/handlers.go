package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	mux  *http.ServeMux
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *http.ServeMux {
	h := &Handler{
		mux:  http.NewServeMux(),
		pool: pool,
	}

	h.mux.HandleFunc("/", homeHandler)
	h.mux.HandleFunc("/db", h.dbHandler)

	return h.mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func (h *Handler) dbHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()
	if err := h.pool.Ping(ctx); err != nil {
		http.Error(w, "db not ready", http.StatusServiceUnavailable)
		return
	}
	w.Write([]byte("ok"))
}
