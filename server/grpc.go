package server

import (
	"base/internal/constants"
	"base/pkg/auth"
	e "base/pkg/grpc/errors"
	"context"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func RunGRPCServer(address string, port int, jwtSecret string, registerServer func(server *grpc.Server)) ServerShutdownFunc {
	addr := fmt.Sprintf("%s:%d", address, port)
	return RunGRPCServerOnAddr(addr, jwtSecret, registerServer)
}

func RunGRPCServerOnAddr(addr string, jwtSecret string, registerServer func(server *grpc.Server)) ServerShutdownFunc {
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			authInterceptor(jwtSecret),
		),
	)
	registerServer(grpcServer)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err = grpcServer.Serve(listen)
		if err != nil {
			log.Fatal("[server] Unable to start GRPC server: ", err)
		}
	}()

	log.Printf("[server] Started listening on: %s", addr)

	return func(ctx context.Context) error {
		grpcServer.GracefulStop()
		return nil
	}
}

func authInterceptor(jwtSecret string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, e.NewGRPCError(fmt.Errorf("failed to read metadata")).
				WithCode(codes.Unauthenticated).
				Err()
		}

		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {
			return nil, e.NewGRPCError(fmt.Errorf("authorization metadata not found")).
				WithCode(codes.Unauthenticated).
				Err()
		}

		reqToken := authHeader[0]
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) < 2 {
			return nil, e.NewGRPCError(fmt.Errorf("missing bearer token")).
				WithCode(codes.Unauthenticated).
				Err()
		}

		jwtToken := splitToken[1]
		jwtManager := auth.NewJWTManager(jwtSecret)

		claims, err := jwtManager.ValidateTokenAndExtractData(jwtToken)
		if err != nil {
			return nil, e.NewGRPCError(fmt.Errorf("token validation failed: %v", err)).
				WithCode(codes.Unauthenticated).
				Err()
		}

		nextCtx := context.WithValue(ctx, constants.ContextKeyUserID, claims.UserID)

		h, err := handler(nextCtx, req)
		return h, err
	}
}
