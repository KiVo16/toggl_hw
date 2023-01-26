package main

import (
	db "base/infrastructure/db/sqlite"
	repo "base/infrastructure/repo"
	"base/internal/app"
	"base/internal/app/handlers"
	ports "base/ports/http"
	"base/server"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	ctx, cancelMainCtx := context.WithCancel(context.Background())
	sqlLite, err := db.NewSQLiteDB(ctx)
	if err != nil {
		panic(err)
	}

	repo := repo.NewRepoWithSQLite(*sqlLite)
	appHandlers := handlers.NewHandlers(repo)
	app := app.NewApp(appHandlers)

	srv := server.RunHTTPServer(func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(
			ports.NewHttpServer(*app),
			router,
		)
	})

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	log.Println("test2")

	ctx, cancelShutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelShutdown()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server gracefull shutdown failed: %v", err)
	}

	cancelMainCtx()

	log.Println("Server is shutting down")
	os.Exit(0)

}
