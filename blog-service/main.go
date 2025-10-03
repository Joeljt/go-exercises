package main

import (
	"net/http"
	"time"

	"github.com/joeljt/blog-service/internal/routers"
)

func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":1219",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
