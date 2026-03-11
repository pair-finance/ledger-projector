package enums

import "iter"

type EventType string

const (
	EventTypeSample = "sample"
)

func (e EventType) IsValid() bool {
	switch e {
	case EventTypeSample:
		return true
	}

	return false
}

func allTypes() []EventType {
	return []EventType{
		EventTypeSample,
		// Add future event types here
	}
}

func IterateEventTypes() iter.Seq[EventType] {
	return func(yield func(EventType) bool) {
		for _, t := range allTypes() {
			if !yield(t) {
				return
			}
		}
	}
}
