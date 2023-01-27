package main

import (
	db "base/infrastructure/db/sqlite"
	repo "base/infrastructure/repo"
	"base/internal/app"
	"base/internal/app/handlers"
	"base/internal/config"
	grpcPorts "base/ports/grpc"
	pb "base/ports/grpc/proto"
	httpPorts "base/ports/http"
	"base/server"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

func main() {
	conf, err := config.LoadAndValidateConfig()
	if err != nil {
		log.Fatal(err)
	}
	dbConf := conf.Database
	apiConf := conf.API
	authConf := conf.Auth

	ctx, cancelMainCtx := context.WithCancel(context.Background())
	sqlLite, err := db.NewSQLiteDB(ctx, db.WithFile(dbConf.File))
	if err != nil {
		log.Fatal(err)
	}

	repo := repo.NewRepoWithSQLite(*sqlLite)
	appHandlers := handlers.NewHandlers(repo)
	app := app.NewApp(appHandlers)

	var serverShutDownFunction server.ServerShutdownFunc

	if apiConf.Mode == config.ModeHttp {
		serverShutDownFunction = server.RunHTTPServer(apiConf.Host, apiConf.Port, authConf.JWTSecret,
			func(router chi.Router) http.Handler {
				return httpPorts.HandlerFromMux(
					httpPorts.NewHttpServer(*app),
					router,
				)
			})
	} else if apiConf.Mode == config.ModeGrpc {
		serverShutDownFunction = server.RunGRPCServer(apiConf.Host, apiConf.Port, authConf.JWTSecret,
			func(server *grpc.Server) {
				svc := grpcPorts.NewGrpcServer(*app)
				pb.RegisterQuestionsServiceServer(server, svc)
			})

	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	ctx, cancelShutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelShutdown()

	if err := serverShutDownFunction(ctx); err != nil {
		log.Fatalf("Server gracefull shutdown failed: %v", err)
	}

	cancelMainCtx()
	time.Sleep(time.Second * 2)
	// added additional 2 seconds to let the database close - it uses the context above that needs to be canceled
	// it can be resolved in a different way - for example by taking a similar approach as closing
	// HTTP server (passing ctx to shutdown function) above

	log.Println("Server is shutting down")
}
