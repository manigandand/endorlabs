package inmem

import (
	"sync"

	"github.com/manigandand/endorlabs/store/adapter"
)

// NewAdapter returns store inmemory adapter(*Client)
func NewAdapter() adapter.ObjectDB {
	c := &Client{
		mu:           sync.RWMutex{},
		objects:      make(map[string]adapter.Object),
		spreadByName: make(map[string]adapter.Object),
		// grpByKind:    make(map[string][]adapter.Object),
	}

	return c
}
