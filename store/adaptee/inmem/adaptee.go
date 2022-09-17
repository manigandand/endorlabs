package inmem

import (
	"github.com/manigandand/endorlabs/store/adapter"
)

// NewAdapter returns store inmemory adapter(*Client)
func NewAdapter() adapter.ObjectDB {
	c := &Client{}

	return c
}
