package enums

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
