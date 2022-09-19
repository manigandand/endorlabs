// Package adapter has the signature of store layer interface
package adapter

import "context"

// Object interface
type Object interface {
	// GetKind returns the type of the object.
	GetKind() string

	// GetID returns a unique UUID for the object.
	GetID() string

	// GetName returns the name of the object. Names are not unique.
	GetName() string

	// SetID sets the ID of the object.
	SetID(string)

	// SetName sets the name of the object.
	SetName(string)
}

// ObjectDB interface
type ObjectDB interface {
	// Store will store the object in the data store. The object will have a
	// name and kind, and the Store method should create a unique ID.
	Store(ctx context.Context, object Object) error

	// GetObjectByID will retrieve the object with the provided ID.
	GetObjectByID(ctx context.Context, id string) (Object, error)

	// GetObjectByName will retrieve the object with the given name.
	GetObjectByName(ctx context.Context, name string) (Object, error)

	// ListObjects will return a list of all objects of the given kind.
	ListObjects(ctx context.Context, kind string) ([]Object, error)

	// DeleteObject will delete the object.
	DeleteObject(ctx context.Context, id string) error

	Close()
}
