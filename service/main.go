package service

import (
	"net"
	"net/http"
	"time"

	"github.com/EricTusk/template-http-grpc/api"
	"github.com/EricTusk/template-http-grpc/worker"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

const (
	network = "tcp"
)

func RunGRPCServer(cfg *worker.Config) error {
	grpc_prometheus.EnableHandlingTimeHistogram()
	grpcLog := logrus.WithField("service", "template-http-grpc")
	opt := grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_prometheus.UnaryServerInterceptor,
		grpc_logrus.UnaryServerInterceptor(grpcLog),
		UnaryServerInterceptor(),
	))

	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(cfg.MaxMsgSize),
		grpc.MaxSendMsgSize(cfg.MaxMsgSize),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     time.Hour,
			MaxConnectionAge:      time.Hour,
			MaxConnectionAgeGrace: time.Minute * 10,
			Time:                  time.Minute * 10,
			Timeout:               time.Second * 20,
		}),
		opt,
	)
	service, err := NewService()
	if err != nil {
		logrus.Errorf("failed to new service with error: %v", err)
		return err
	}
	api.RegisterTemplateHTTPGRPCServiceServer(grpcServer, service)
	reflection.Register(grpcServer)

	grpcEndpoint := cfg.GRPCEndpoint
	listener, err := net.Listen(network, grpcEndpoint)
	if err != nil {
		logrus.Errorf("failed to listen %s with error: %v", grpcEndpoint, err)
		return err
	}
	defer func() {
		if err := listener.Close(); err != nil {
			logrus.Errorf("failed to close %s with error %v", grpcEndpoint, err)
		}
		// grpcServer.GracefulStop()
	}()

	// Run grpc server
	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatalf("failed to serve with error: %v", err)
		return err
	}

	return nil
}

func RunHTTPServer(cfg *worker.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))

	dialOptions := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(cfg.MaxMsgSize), grpc.MaxCallSendMsgSize(cfg.MaxMsgSize)),
	}
	if err := api.RegisterTemplateHTTPGRPCServiceHandlerFromEndpoint(ctx, mux, cfg.GRPCEndpoint, dialOptions); err != nil {
		logrus.Fatalf("failed to register handler with error: %v", err)
	}

	http.Handle("/", mux)
	/*	fhandler := fasthttpadaptor.NewFastHTTPHandler(mux)
		if err:= fasthttp.ListenAndServe(httpEndpoint, fhandler); err!=nil {
			logrus.Fatal("RunHTTPServer(). start HTTP server failed. ", err)
		}*/

	if err := http.ListenAndServe(cfg.HTTPEndpoint, nil); err != nil {
		logrus.Fatalf("failed to listen with error: %v", err)
	}
}

func RunMetricsServer(cfg *worker.Config) {
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(cfg.MetricsEndpoint, nil); err != nil {
		logrus.Errorf("failed to run metrics server with error: %v", err)
	}
}
