package sqlcomposers

import "github.com/pair-finance/ledger-projector/internal/enums"

type Composer func() // type describing all sql compose funcs

func GetComposersMapped() map[enums.EventType]Composer {
	return map[enums.EventType]Composer{
		enums.EventTypeSample: ComposeSampleEventQuery,
	}
}
