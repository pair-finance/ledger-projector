package event

import "github.com/pair-finance/ledger-projector/internal/enums"

type HandlerFunc func()

type Service struct {
	batches map[enums.EventType]HandlerFunc
}
