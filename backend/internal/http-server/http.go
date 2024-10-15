package httpServer

import (
	"net/http"
	"time"
)

func NewServer(address string, router http.Handler, timeout time.Duration, IdleTimeout time.Duration) *http.Server {
	return &http.Server{
		Addr:         address,
		Handler:      router,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		IdleTimeout:  IdleTimeout,
	}
}
