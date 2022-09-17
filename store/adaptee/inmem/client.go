// Package inmem implements the store adapter for inmemory data store
package inmem

import (
	"context"

	"github.com/manigandand/endorlabs/store/adapter"
)

// Client struct holds the store connection
type Client struct{}

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
