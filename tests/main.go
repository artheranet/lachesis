package main

import (
	"fmt"
	"io"

	"github.com/artheranet/lachesis/abft"
	"github.com/artheranet/lachesis/inter/idx"
	"github.com/artheranet/lachesis/kvdb"
	"github.com/artheranet/lachesis/kvdb/memorydb"
	"github.com/artheranet/lachesis/utils/adapters"
	"github.com/artheranet/lachesis/vecfc"
)

func main() {
	openEDB := func(epoch idx.Epoch) kvdb.Store {
		return memorydb.New()
	}

	crit := func(err error) {
		panic(err)
	}

	store := abft.NewStore(memorydb.New(), openEDB, crit, abft.LiteStoreConfig())
	restored := abft.NewIndexedLachesis(store, nil, &adapters.VectorToDagIndexer{Index: vecfc.NewIndex(crit, vecfc.LiteConfig())}, crit, abft.LiteConfig())

	// prevent compiler optimizations
	fmt.Fprint(io.Discard, restored == nil)
}
