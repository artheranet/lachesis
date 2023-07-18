package tdag

import (
	"github.com/artheranet/lachesis/hash"
	"github.com/artheranet/lachesis/inter/dag"
)

type TestEvent struct {
	dag.MutableBaseEvent
	Name string
}

func (e *TestEvent) AddParent(id hash.Event) {
	parents := e.Parents()
	parents.Add(id)
	e.SetParents(parents)
}
