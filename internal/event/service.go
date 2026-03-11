package event

import (
	"context"
	"log/slog"

	"github.com/pair-finance/ledger-projector/internal/enums"
	"github.com/pair-finance/ledger-projector/internal/event/batcher"
	"github.com/pair-finance/ledger-projector/internal/event/sqlcomposers"
)

type Service struct {
	composers map[enums.EventType]sqlcomposers.Composer
	batches   map[enums.EventType]*batcher.Batch
	logger    *slog.Logger
}

func NewService(logger *slog.Logger) *Service {
	batches := make(map[enums.EventType]*batcher.Batch, 0)
	for et := range enums.IterateEventTypes() {
		batches[et] = batcher.NewEventsBatch()
	}

	return &Service{
		composers: sqlcomposers.GetComposersMapped(),
		batches:   batches,
		logger:    logger,
	}
}

func (s *Service) Run(ctx context.Context) error {
	return nil
}
