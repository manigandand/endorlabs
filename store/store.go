// Package store implements db layer of different store adapters
package store

import (
	"log"

	"github.com/manigandand/endorlabs/store/adaptee/inmem"
	"github.com/manigandand/endorlabs/store/adaptee/sqlite"
	"github.com/manigandand/endorlabs/store/adapter"
)

// Init loads the sample data and prepares the store layer
func Init(dbType string) adapter.ObjectDB {
	var store adapter.ObjectDB

	// store inmemory adapter ...
	switch dbType {
	case "inmemory":
		store = inmem.NewAdapter()
	case "sqlite":
		store = sqlite.NewAdapter("")
	}

	if store == nil {
		log.Fatalf("🦠store initialize failed 👎")
	}
	log.Println("Inited store...👍")
	return store
}
