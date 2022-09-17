package sqlite

import (
	"github.com/manigandand/endorlabs/store/adapter"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewAdapter returns store inmemory adapter(*Client)
func NewAdapter(dbName string) adapter.ObjectDB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	c := &Client{
		db: db,
	}

	return c
}
