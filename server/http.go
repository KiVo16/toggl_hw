package server

import (
	"base/internal/constants"
	"base/pkg/auth"
	e "base/pkg/http/errors"
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func RunHTTPServer(address string, port int, jwtSecret string, createHandler func(router chi.Router) http.Handler) ServerShutdownFunc {
	addr := fmt.Sprintf("%s:%d", address, port)
	return RunHTTPServerOnAddr(addr, jwtSecret, createHandler)
}

func RunHTTPServerOnAddr(addr, jwtSecret string, createHandler func(router chi.Router) http.Handler) ServerShutdownFunc {
	apiRouter := chi.NewRouter()
	setHTTPMiddlewares(apiRouter, jwtSecret)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api/v1", createHandler(apiRouter))

	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      rootRouter,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("[server] Unable to start HTTP server: ", err)
		}
	}()

	log.Printf("[server] Started listening on: %s", addr)

	return func(ctx context.Context) error {
		return server.Shutdown(ctx)
	}
}

func setHTTPMiddlewares(router *chi.Mux, jwtSecret string) {
	router.Use(middleware.Logger)
	router.Use(authMiddleware(jwtSecret))
}

func authMiddleware(jwtSecret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			reqToken := r.Header.Get("Authorization")
			splitToken := strings.Split(reqToken, "Bearer ")
			if len(splitToken) < 2 {
				e.NewHttpError(fmt.Errorf("missing bearer token")).
					WithCode(http.StatusUnauthorized).
					Handle(w)
				return
			}

			jwtToken := splitToken[1]
			jwtManager := auth.NewJWTManager(jwtSecret)

			claims, err := jwtManager.ValidateTokenAndExtractData(jwtToken)
			if err != nil {
				e.NewHttpError(fmt.Errorf("token validation failed: %v", err)).
					WithCode(http.StatusUnauthorized).
					Handle(w)
				return
			}

			ctx := context.WithValue(r.Context(), constants.ContextKeyUserID, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
