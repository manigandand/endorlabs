// Package inmem implements the store adapter for inmemory data store
package inmem

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/manigandand/endorlabs/store/adapter"
)

// Client struct holds the store connection
type Client struct {
	mu sync.RWMutex

	// objects stored by id
	objects map[string]adapter.Object
	// objects stored by name, to search by name
	spreadByName map[string]adapter.Object
	// objects grouped by kind
	// grpByKind    map[string][]adapter.Object
}

// Store will store the object in the data store. The object will have a
// name and kind, and the Store method should create a unique ID.
func (c *Client) Store(ctx context.Context, object adapter.Object) error {
	kind := object.GetKind()
	if kind == "" {
		return errors.New("invalid object kind")
	}
	id := uuid.NewString()
	// set id
	object.SetID(id)
	if object.GetID() != id {
		return errors.New("invalid object: id not set")
	}

	c.mu.Lock()
	c.objects[id] = object
	if object.GetName() != "" {
		c.spreadByName[object.GetName()] = object
	}
	c.mu.Unlock()

	return nil
}

// GetObjectByID will retrieve the object with the provided ID.
func (c *Client) GetObjectByID(ctx context.Context, id string) (adapter.Object, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	obj, ok := c.objects[id]
	if !ok {
		return nil, errors.New("object not found")
	}

	return obj, nil
}

// GetObjectByName will retrieve the object with the given name.
func (c *Client) GetObjectByName(ctx context.Context, name string) (adapter.Object, error) {
	if name == "" {
		return nil, errors.New("invalid name")
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	obj, ok := c.spreadByName[name]
	if !ok {
		return nil, errors.New("object not found")
	}

	return obj, nil
}

// ListObjects will return a list of all objects of the given kind.
func (c *Client) ListObjects(ctx context.Context, kind string) ([]adapter.Object, error) {
	if kind == "" {
		return nil, errors.New("invalid kind")
	}

	c.mu.RLock()
	defer c.mu.RUnlock()

	var objs []adapter.Object
	for _, obj := range c.objects {
		if obj.GetKind() == kind {
			objs = append(objs, obj)
		}
	}

	return objs, nil
}

// DeleteObject will delete the object.
func (c *Client) DeleteObject(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("invalid id")
	}

	c.mu.Lock()

	obj, ok := c.objects[id]
	if !ok {
		return errors.New("object not found")
	}

	delete(c.objects, id)
	delete(c.spreadByName, obj.GetName())

	c.mu.Unlock()

	return nil
}
