// Package sqlite implements the store adapter for sqlite data store
package sqlite

import (
	"context"

	"github.com/manigandand/endorlabs/store/adapter"
	"gorm.io/gorm"
)

// Client struct holds the store connection
type Client struct {
	db *gorm.DB
}

// Store will store the object in the data store. The object will have a
// name and kind, and the Store method should create a unique ID.
func (c *Client) Store(ctx context.Context, object adapter.Object) error {
	return nil
}

// GetObjectByID will retrieve the object with the provided ID.
func (c *Client) GetObjectByID(ctx context.Context, id string) (adapter.Object, error) {
	return nil, nil
}

// GetObjectByName will retrieve the object with the given name.
func (c *Client) GetObjectByName(ctx context.Context, name string) (adapter.Object, error) {
	return nil, nil
}

// ListObjects will return a list of all objects of the given kind.
func (c *Client) ListObjects(ctx context.Context, kind string) ([]adapter.Object, error) {
	return nil, nil
}

// DeleteObject will delete the object.
func (c *Client) DeleteObject(ctx context.Context, id string) error {
	return nil
}
