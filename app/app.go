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

	handlers "github.com/adamlahbib/postgresser/api/handlers"
	"github.com/adamlahbib/postgresser/api/middlewares"
	pb "github.com/adamlahbib/postgresser/api/proto"
	"github.com/adamlahbib/postgresser/api/servers"
	"github.com/adamlahbib/postgresser/services"
	prometheusService "github.com/adamlahbib/postgresser/services"
	grpcLogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpcPrometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
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

	// Prometheus Metrics Services
	serverMetrics := grpcPrometheus.NewServerMetrics(
		grpcPrometheus.WithServerHandlingTimeHistogram(),
	)
	registry := prometheus.NewRegistry() // create a new registry for the metrics to be registered
	registry.MustRegister(serverMetrics) // register the metrics to the registry
	// Prometheus custom metrics
	customMetrics, err := prometheusService.NewPrometheusService(prometheusService.GetMetricsDefinition())
	if err != nil {
		log.Fatalf("failed to create prometheus service: %v", err)
	}
	registry.MustRegister(customMetrics)

	// Postgres service init
	postgresService := services.NewPostgres(kubeClient)

	sCtx := serverContext(context.Background())

	healthcheckHandlers := handlers.NewHealthcheck(registry)
	httpServer := servers.NewHealthcheck("", "8080", 10, *healthcheckHandlers)
	httpServer.Start()

	// grpc server init
	port := 5000
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("starting gRPC server")
	grpcServer := grpc.NewServer([]grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			middlewares.AuthInterceptor,
			grpcLogrus.UnaryServerInterceptor(logrus.NewEntry(logrus.New()), []grpcLogrus.Option{}...),
			serverMetrics.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(),
		),
		grpc.ChainStreamInterceptor(
			serverMetrics.StreamServerInterceptor(),
			grpcRecovery.StreamServerInterceptor(),
		),
	}...)
	pb.RegisterPostgresServiceServer(grpcServer, servers.NewPostgres(postgresService))

	go func() {
		log.Printf("gRPC server listening on port: %d", listener.Addr())
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	<-sCtx.Done()
	registry.Unregister(customMetrics)
	registry.Unregister(serverMetrics)
	grpcServer.GracefulStop()
	if err := httpServer.Stop(); err != nil {
		log.Fatalf("failed to stop healthcheck server: %v", err)
	}
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
