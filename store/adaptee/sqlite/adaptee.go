package sqlite

import (
	"encoding/json"
	"log"
	"time"

	"github.com/manigandand/endorlabs/store/adapter"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Object reference
type Object struct {
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Name string          `json:"name" gorm:"not null;uniqueIndex:idx_name_kind,where:deleted_at IS NULL"`
	Kind string          `json:"kind" gorm:"not null;uniqueIndex:idx_name_kind,where:deleted_at IS NULL;index:idx_kind"`
	Obj  json.RawMessage `json:"obj" gorm:"not null"`
}

// NewAdapter returns store inmemory adapter(*Client)
func NewAdapter(dbName string) adapter.ObjectDB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	c := &Client{
		db: db,
	}

	if err := db.AutoMigrate(&Object{}); err != nil {
		log.Printf("Failed to set up DB table(objects) with error: %s\n", err.Error())
	}

	return c
}

// Close the connections
func Close(c *Client) {
	c.db.Migrator().DropTable(&Object{})

	db, _ := c.db.DB()
	db.Close()
}
