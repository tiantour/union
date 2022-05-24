package cache

import (
	"log"

	"github.com/dgraph-io/ristretto"
)

var cache *ristretto.Cache

func init() {
	var err error
	cache, err = ristretto.NewCache(&ristretto.Config{
		NumCounters:        1e7,     // number of keys to track frequency of (10M).
		MaxCost:            1 << 30, // maximum cost of cache (1GB).
		BufferItems:        64,      // number of keys per Get buffer.
		IgnoreInternalCost: true,
	})
	if err != nil {
		log.Fatal(err)
	}
}
