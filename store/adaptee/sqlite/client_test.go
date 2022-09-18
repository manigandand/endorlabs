package sqlite

import (
	"context"
	"testing"
	"time"

	"github.com/manigandand/endorlabs/schema"
	"github.com/manigandand/endorlabs/store/adapter"
	"github.com/stretchr/testify/assert"
)

func tearDown(client adapter.ObjectDB) {
	c, ok := client.(*Client)
	if !ok {
		return
	}

	c.db.Delete(&Object{}, "deleted_at IS NULL")
}

func tearDownDB(client adapter.ObjectDB) {
	c, ok := client.(*Client)
	if !ok {
		return
	}

	c.db.Migrator().DropTable(&Object{})

	db, _ := c.db.DB()
	db.Close()
}

func TestSqliteStore(t *testing.T) {
	bg := context.Background()
	client := NewAdapter("endorlabs-dev.db")
	defer tearDownDB(client)

	t.Run("Save method with invalid objects - no kind", func(t *testing.T) {
		dum := &dummyObject{}

		// save
		err := client.Store(bg, dum)
		assert.Equal(t, "invalid object kind", err.Error())
	})

	t.Run("Save method with invalid objects - no id", func(t *testing.T) {
		dum := &dummyObject2{}

		// save
		err := client.Store(bg, dum)
		assert.Equal(t, "invalid object: id not set", err.Error())
	})

	t.Run("Save method", func(t *testing.T) {
		bt, err := time.Parse("2006-01-02T15:04:05.000Z", "1993-07-22T05:04:05.000Z")
		assert.Equal(t, nil, err)

		person := &schema.Person{
			Name:      "Manigandan Dharmalingam",
			LastName:  "Dharmalingam",
			Birthday:  "22/07/1993",
			BirthDate: bt,
		}

		assert.Equal(t, "", person.GetID())

		// save
		err = client.Store(bg, person)
		assert.Equal(t, nil, err)

		// id should have updtaed
		assert.NotEqual(t, "", person.GetID())

		tearDown(client)
	})

	t.Run("GetObjectByID method - invalid id", func(t *testing.T) {
		obj, err := client.GetObjectByID(bg, "")
		assert.Equal(t, "invalid id", err.Error())
		assert.Equal(t, nil, obj)

		obj, err = client.GetObjectByID(bg, "dummy-id")
		assert.Equal(t, "object not found", err.Error())
		assert.Equal(t, nil, obj)

		tearDown(client)
	})

	t.Run("GetObjectByID method", func(t *testing.T) {
		bt, err := time.Parse("2006-01-02T15:04:05.000Z", "1993-07-22T05:04:05.000Z")
		assert.Equal(t, nil, err)

		person := &schema.Person{
			Name:      "Manigandan2",
			LastName:  "Dharmalingam",
			Birthday:  "22/07/1993",
			BirthDate: bt,
		}
		assert.Equal(t, "", person.GetID())

		// save
		err = client.Store(bg, person)
		assert.Equal(t, nil, err)
		// id should have updtaed
		assert.NotEqual(t, "", person.GetID())

		obj, err := client.GetObjectByID(bg, person.GetID())
		assert.Equal(t, nil, err)

		assert.Equal(t, obj.GetID(), person.GetID())
		assert.Equal(t, obj.GetKind(), person.GetKind())
		assert.Equal(t, obj.GetName(), person.GetName())

		pobj, ok := obj.(*schema.Person)
		assert.Equal(t, true, ok)
		assert.Equal(t, person.Birthday, pobj.Birthday)

		tearDown(client)
	})

	t.Run("GetObjectByName method - invalid name", func(t *testing.T) {
		obj, err := client.GetObjectByName(bg, "")
		assert.Equal(t, "invalid name", err.Error())
		assert.Equal(t, nil, obj)

		obj, err = client.GetObjectByName(bg, "dummy-name")
		assert.Equal(t, "object not found", err.Error())
		assert.Equal(t, nil, obj)
		tearDown(client)
	})

	t.Run("GetObjectByName method", func(t *testing.T) {
		bt, err := time.Parse("2006-01-02T15:04:05.000Z", "1993-07-22T05:04:05.000Z")
		assert.Equal(t, nil, err)

		person := &schema.Person{
			Name:      "Manigandan Dharmalingam",
			LastName:  "Dharmalingam",
			Birthday:  "22/07/1993",
			BirthDate: bt,
		}
		assert.Equal(t, "", person.GetID())

		// save
		err = client.Store(bg, person)
		assert.Equal(t, nil, err)
		// id should have updtaed
		assert.NotEqual(t, "", person.GetID())

		obj, err := client.GetObjectByName(bg, person.GetName())
		assert.Equal(t, nil, err)
		assert.Equal(t, obj.GetID(), person.GetID())
		assert.Equal(t, obj.GetKind(), person.GetKind())
		assert.Equal(t, obj.GetName(), person.GetName())

		pobj, ok := obj.(*schema.Person)
		assert.Equal(t, true, ok)
		assert.Equal(t, person.Birthday, pobj.Birthday)

		tearDown(client)
	})

	t.Run("ListObjects method - invalid name", func(t *testing.T) {
		obj, err := client.ListObjects(bg, "")
		assert.Equal(t, "invalid kind", err.Error())
		assert.Equal(t, 0, len(obj))

		obj, err = client.ListObjects(bg, "dummy-kind")
		assert.Equal(t, nil, err)
		assert.Equal(t, 0, len(obj))

		tearDown(client)
	})

	t.Run("ListObjects method", func(t *testing.T) {
		bt, err := time.Parse("2006-01-02T15:04:05.000Z", "1993-07-22T05:04:05.000Z")
		assert.Equal(t, nil, err)

		person := &schema.Person{
			Name:      "Manigandan Dharmalingam",
			LastName:  "Dharmalingam",
			Birthday:  "22/07/1993",
			BirthDate: bt,
		}
		assert.Equal(t, "", person.GetID())

		// save
		err = client.Store(bg, person)
		assert.Equal(t, nil, err)
		// id should have updtaed
		assert.NotEqual(t, "", person.GetID())

		obj, err := client.ListObjects(bg, person.GetKind())
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, len(obj))

		assert.Equal(t, obj[0].GetID(), person.GetID())
		assert.Equal(t, obj[0].GetKind(), person.GetKind())
		assert.Equal(t, obj[0].GetName(), person.GetName())

		pobj, ok := obj[0].(*schema.Person)
		assert.Equal(t, true, ok)
		assert.Equal(t, person.Birthday, pobj.Birthday)
		tearDown(client)
	})

	t.Run("DeleteObject method - invalid id", func(t *testing.T) {
		err := client.DeleteObject(bg, "")
		assert.Equal(t, "invalid id", err.Error())
	})

	t.Run("DeleteObject method", func(t *testing.T) {
		bt, err := time.Parse("2006-01-02T15:04:05.000Z", "1993-07-22T05:04:05.000Z")
		assert.Equal(t, nil, err)

		person := &schema.Person{
			Name:      "Manigandan Dharmalingam",
			LastName:  "Dharmalingam",
			Birthday:  "22/07/1993",
			BirthDate: bt,
		}
		assert.Equal(t, "", person.GetID())

		// save
		err = client.Store(bg, person)
		assert.Equal(t, nil, err)
		// id should have updtaed
		assert.NotEqual(t, "", person.GetID())

		err = client.DeleteObject(bg, person.GetID())
		assert.Equal(t, nil, err)

		obj, err := client.GetObjectByID(bg, person.GetID())
		assert.Equal(t, "object not found", err.Error())
		assert.Equal(t, nil, obj)
		tearDown(client)
	})
}

type dummyObject struct {
	ID   string
	Name string
}

func (p *dummyObject) GetKind() string {
	return ""
}

func (p *dummyObject) GetID() string {
	return p.ID
}

func (p *dummyObject) GetName() string {
	return p.Name
}

func (p *dummyObject) SetID(s string) {
	p.ID = s
}

func (p *dummyObject) SetName(s string) {
	p.Name = s
}

type dummyObject2 struct {
	ID   string
	Name string
}

func (p *dummyObject2) GetKind() string {
	return "dummyObject2"
}

func (p *dummyObject2) GetID() string {
	return p.ID
}

func (p *dummyObject2) GetName() string {
	return p.Name
}

func (p *dummyObject2) SetID(s string) {
	// no imple
}

func (p *dummyObject2) SetName(s string) {
	p.Name = s
}
