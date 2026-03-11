package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pair-finance/ledger-projector/internal/event"
	"golang.org/x/sync/errgroup"
)

const tick = 500 * time.Millisecond

func main() {
	ctx := context.Background()

	logger := slog.Default()

	eventService := event.NewService(logger)

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	sigs := make(chan os.Signal, 1)
	g, _ := errgroup.WithContext(ctx)
	g.Go(
		func() error {
			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			return nil
		},
	)

	logger.Info("Ledger Projector started.")

	for {
		select {
		case <-ticker.C:
			eventService.Run(ctx)
		case <-sigs:
			logger.WarnContext(ctx, "Terminated by SIGINT/SIGTERM.")
			return
		}
	}
}
