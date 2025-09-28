package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/vivek869858/GoLiveProject/chat/foundation/logger"
)

func main() {
	traceIdfn := func(ctx context.Context) string { return "trace-123" } // Example trace ID function
	log := logger.New(os.Stdout, logger.LevelInfo, "CAP", traceIdfn)

	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.Error(ctx, "Startup", "err", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *logger.Logger) error {

	log.Info(ctx, "Running the application", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	log.Info(ctx, "Running the application", "status", "started")
	defer log.Info(ctx, "Running the application", "GOMAXPROCS", "shuting down")

	shutDown := make(chan os.Signal, 1)
	signal.Notify(shutDown, syscall.SIGINT, syscall.SIGTERM)
	<-shutDown

	return nil
}
