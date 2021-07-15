package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/EricTusk/template-http-grpc/service"
	"github.com/EricTusk/template-http-grpc/version"
	"github.com/EricTusk/template-http-grpc/worker"
	"github.com/onrik/logrus/filename"
	"github.com/sirupsen/logrus"
)

var (
	printVersion = flag.Bool("version", false, "print version of this build")
	verbose      = flag.Bool("verbose", false, "verbose output")
	configFile   = flag.String("config", "example/config.json", "config file path")
)

func main() {
	flag.Parse()
	if *printVersion {
		version.PrintFullVersionInfo()
		return
	}

	logrus.AddHook(filename.NewHook())
	formatter := new(logrus.TextFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	formatter.FullTimestamp = true
	logrus.SetFormatter(formatter)

	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	cfg, err := worker.LoadConfig(*configFile)
	if err != nil {
		logrus.Fatalf("failed to load config from file %s with error: %v", *configFile, err)
	}

	worker, err := worker.NewWorker(cfg)
	if err != nil {
		logrus.Fatalf("failed to create worker with error: %v", err)
	}

	go func() {
		if err := service.RunGRPCServer(cfg); err != nil {
			logrus.Fatalf("failed to run grpc server with error: %v", err)
		}
	}()

	go service.RunHTTPServer(cfg)

	go service.RunMetricsServer(cfg)

	if err := worker.Run(); err != nil {
		logrus.Fatalf("failed to start worker with error: %v", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	sig := <-sigs
	logrus.Infof("received signal %v", sig)

	worker.Stop()
}
