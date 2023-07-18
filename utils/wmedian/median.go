package wmedian

import (
	"github.com/artheranet/lachesis/inter/pos"
)

type WeightedValue interface {
	Weight() pos.Weight
}

func Of(values []WeightedValue, stop pos.Weight) WeightedValue {
	// Calculate weighted median
	var curWeight pos.Weight
	for _, value := range values {
		curWeight += value.Weight()
		if curWeight >= stop {
			return value
		}
	}
	panic("invalid median")
}
