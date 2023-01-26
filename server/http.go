package server

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func RunHTTPServer(createHandler func(router chi.Router) http.Handler) *http.Server {
	return RunHTTPServerOnAddr("0.0.0.0:8001", createHandler)
}

func RunHTTPServerOnAddr(addr string, createHandler func(router chi.Router) http.Handler) *http.Server {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api/v1", createHandler(apiRouter))

	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      rootRouter,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Unable to start HTTP server: ", err)
	}

	return server
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.Logger)
}
