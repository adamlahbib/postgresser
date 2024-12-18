package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	pb "github.com/adamlahbib/postgresser/api/proto"
	"github.com/adamlahbib/postgresser/api/servers"
	"github.com/adamlahbib/postgresser/services"
	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Start() {
	// kubernetes init
	kubeConfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatalf("failed to load kubeConfig path: %v", err)
	}
	kubeClient, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalf("failed to load kubeConfig: %v", err)
	}
	// Postgres service init
	postgresService := services.NewPostgres(kubeClient)

	// grpc server init
	port := 5000
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("starting gRPC server")
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	pb.RegisterPostgresServiceServer(grpcServer, servers.NewPostgres(postgresService))
	sCtx := serverContext(context.Background())
	go func() {
		log.Printf("gRPC server listening on port: %d", listener.Addr())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	<-sCtx.Done()
	grpcServer.GracefulStop()
	log.Println("gRPC server stopped")
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		s := <-c
		log.Printf("received signal: %v", s)
		cancel()
	}()
	return ctx
}
