// Package sqlite implements the store adapter for sqlite data store
package sqlite

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/manigandand/endorlabs/schema"
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

	objRaw, err := json.Marshal(&object)
	if err != nil {
		return err
	}

	newObj := &Object{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      object.GetName(),
		Kind:      object.GetKind(),
		Obj:       objRaw,
	}

	// save
	return c.db.Create(newObj).Error
}

// GetObjectByID will retrieve the object with the provided ID.
func (c *Client) GetObjectByID(ctx context.Context, id string) (adapter.Object, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}

	var res Object
	if err := c.db.First(&res).Where("id = ?", id).Error; err != nil {
		return nil, errors.New("object not found")
	}

	return decodeObject(&res)
}

func decodeObject(obj *Object) (adapter.Object, error) {
	switch obj.Kind {
	case "*schema.Person":
		var v schema.Person
		if err := json.Unmarshal(obj.Obj, &v); err != nil {
			return nil, err
		}

		return &v, nil
	case "*schema.Animal":
		var v schema.Animal
		if err := json.Unmarshal(obj.Obj, &v); err != nil {
			return nil, err
		}

		return &v, nil
	}
	return nil, errors.New("decode: invalid kind")
}

// GetObjectByName will retrieve the object with the given name.
func (c *Client) GetObjectByName(ctx context.Context, name string) (adapter.Object, error) {
	if name == "" {
		return nil, errors.New("invalid name")
	}

	var res Object
	if err := c.db.First(&res).Where("name = ?", name).Error; err != nil {
		return nil, errors.New("object not found")
	}

	return decodeObject(&res)
}

// ListObjects will return a list of all objects of the given kind.
func (c *Client) ListObjects(ctx context.Context, kind string) ([]adapter.Object, error) {
	if kind == "" {
		return nil, errors.New("invalid kind")
	}

	var results []Object
	if err := c.db.Find(&results, "kind = ?", kind).Error; err != nil {
		return nil, errors.New("object not found")
	}

	var objects []adapter.Object
	for _, res := range results {
		obj, err := decodeObject(&res)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}

	return objects, nil
}

// DeleteObject will delete the object.
func (c *Client) DeleteObject(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("invalid id")
	}

	if err := c.db.Delete(&Object{}, "id = ? AND deleted_at IS NULL", id).Error; err != nil {
		return errors.New("object not found " + err.Error())
	}

	return nil
}
