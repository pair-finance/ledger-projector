package batcher

import (
	"fmt"
	"iter"

	"github.com/google/uuid"
	"github.com/pair-finance/ledger-projector/internal/enums"
)

type Batch struct {
	List []Event // make an array and return new batch in AddEvents if over?
	Size int

	typ *enums.EventType // for validation
}

type Event struct {
	ID   uuid.UUID
	Type enums.EventType
}

func NewEventsBatch() *Batch {
	return &Batch{
		List: nil,
		Size: 0,
	}
}

func (b *Batch) AddEvents(events ...Event) error {
	if len(events) < 1 {
		return nil
	}

	if b.typ == nil {
		b.typ = &events[0].Type
	}

	for _, ev := range events {
		if ev.Type != *b.typ {
			return fmt.Errorf("batch is of type [%v], [%v] provided to add", b.typ, ev.Type)
		}
	}

	b.List = append(b.List, events...)
	b.Size += len(events)

	return nil
}

func (b *Batch) Iterate() iter.Seq[Event] {
	return func(yield func(Event) bool) {
		for _, event := range b.List {
			if !yield(event) {
				return
			}
		}
	}
}
