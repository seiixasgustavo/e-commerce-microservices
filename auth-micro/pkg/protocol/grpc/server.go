package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	"github.com/jinzhu/gorm"
	v1 "github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/api/v1"
	"github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/logger"
	"github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/protocol/grpc/middleware"
	service "github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/service/v1"
	"google.golang.org/grpc"
)

// RunServer returns an instance of the gRPC server with the services already linked.
func RunServer(ctx context.Context, port string, db *gorm.DB) error {
	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return nil
	}

	opts := []grpc.ServerOption{}

	opts = middleware.AddLogging(logger.Log, opts)

	server := grpc.NewServer(opts...)
	v1.RegisterAuthServer(server, service.NewAuthServer(db))
	v1.RegisterUserServer(server, service.NewUserServer(db))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for range c {
			logger.Log.Warn("Shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	logger.Log.Info("Starting gRPC server...")
	return server.Serve(lis)
}
