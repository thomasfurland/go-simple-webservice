package httpserver

import (
	"net/http"
	"time"
)

type Options struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}

func New(addr string, h http.Handler, opts Options) *http.Server {
	return &http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  opts.ReadTimeout,
		WriteTimeout: opts.WriteTimeout,
		IdleTimeout:  opts.IdleTimeout,
	}
}
