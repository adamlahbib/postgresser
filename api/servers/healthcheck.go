package servers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/adamlahbib/postgresser/api/handlers"
	"github.com/gorilla/mux"
)

const (
	// default k8s readiness probe path
	defaultReadyProbe = "/.well-known/ready"
	// default k8s liveness probe path
	defaultLiveProbe = "/.well-known/live"
	// default prometheus metrics scraping endpoint path
	defaultMetricsPath = "/metrics"
)

type HealthcheckServer struct {
	httpAddr string
	router   *mux.Router
	timeout  time.Duration
	httpSvr  *http.Server
}

func NewHealthcheck(host, port string, timeout time.Duration, healthcheckHandlers handlers.Healthcheck) HealthcheckServer {
	server := HealthcheckServer{
		httpAddr: host + ":" + port,
		router:   mux.NewRouter(),
		timeout:  timeout,
	}
	server.router.HandleFunc(defaultReadyProbe, healthcheckHandlers.ReadinessProbe).Methods(http.MethodGet)
	server.router.HandleFunc(defaultLiveProbe, healthcheckHandlers.LivenessProbe).Methods(http.MethodGet)
	server.router.HandleFunc(defaultMetricsPath, healthcheckHandlers.MetricsHandler).Methods(http.MethodGet)
	server.httpSvr = &http.Server{
		Addr:              server.httpAddr,
		Handler:           server.router,
		ReadHeaderTimeout: server.timeout,
		ReadTimeout:       server.timeout,
		WriteTimeout:      server.timeout,
	}
	return server
}

func (s *HealthcheckServer) Start() {
	log.Println("starting healthcheck server on ", s.httpAddr)
	go func() {
		if err := s.httpSvr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start healthcheck server: %v", err)
		}
	}()
}

func (s *HealthcheckServer) Stop() error {
	log.Println("stopping healthcheck server")
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()
	if err := s.httpSvr.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
