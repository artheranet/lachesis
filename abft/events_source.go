package abft

import (
	"github.com/artheranet/lachesis/hash"
	"github.com/artheranet/lachesis/inter/dag"
)

// EventSource is a callback for getting events from an external storage.
type EventSource interface {
	HasEvent(hash.Event) bool
	GetEvent(hash.Event) dag.Event
}
