// Package store implements db layer of different store adapters
package store

import (
	"log"

	"github.com/manigandand/endorlabs/store/adaptee/inmem"
	"github.com/manigandand/endorlabs/store/adapter"
)

// Store global store connection interface
var Store adapter.ObjectDB

// Init loads the sample data and prepares the store layer
func Init(dbType string) {
	// store inmemory adapter ...
	switch dbType {
	case "inmemory":
		Store = inmem.NewAdapter()
	case "sqlite":
		// Store = psql.NewAdapter()
	case "cloudsqlpostgres":
	}
	if Store == nil {
		log.Fatalf("🦠store initialize failed 👎")
	}
	log.Println("Inited Store...👍")
}
