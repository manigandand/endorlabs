## Store Layer:

Implements the `ObjectDB` store interface. This implements adapter pattern.

> Directory struct

- schema
  -- containts the different object types which implements the `Object` interface.
- store
  - store.go
    -- inits the different db layer of choice
  - adapter
    -- db interface: `ObjectDB` store interface declarations.
  - adaptee
    -- different adaptees which implements `ObjectDB` store interface.
    - inmem
      -- inmemory store layer
    - sqlite
      -- sqlite store layer

#### How to run

```sh
make test

# run with the exsiting object types
make run
```

> Add new type which implements Object interface

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/manigandand/endorlabs/store"
)

// Bird implements Object interface
type Bird struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Family string `json:"family"`
}

// GetKind returns the object kind
func (p *Bird) GetKind() string {
	return reflect.TypeOf(p).String()
}

// GetID returns the id of the object
func (p *Bird) GetID() string {
	return p.ID
}

// GetName returns the name of the object
func (p *Bird) GetName() string {
	return p.Name
}

// SetID replace the object id
func (p *Bird) SetID(s string) {
	p.ID = s
}

// SetName replace the object name
func (p *Bird) SetName(s string) {
	p.Name = s
}

func main() {
	db := store.Init("inmemory")

	bird := &schema.Bird{
		Name:   "Peacock",
		Family: "Phasianidae",
	}

	if err := db.Store(bg, bird); err != nil {
		log.Fatal("cant store bird 1: ", err.Error())
	}

	birdRes, err := db.GetObjectByID(bg, bird.ID)
	if err != nil {
		log.Fatal("cant get bird 1 by id: ", err.Error())
	}
	fmt.Println(birdRes)

	birdRes, err = db.GetObjectByName(bg, bird.Name)
	if err != nil {
		log.Fatal("cant get bird 1 by name: ", err.Error())
	}
	fmt.Println(birdRes)

	bird2 := &schema.Bird{
		Name:   "Owl",
		Family: "Strigiformes",
	}

	if err := db.Store(bg, bird2); err != nil {
		log.Fatal("cant store bird 2: ", err.Error())
	}
	objs, err = db.ListObjects(bg, "*schema.Bird")
	if err != nil {
		log.Fatal("cant get objects bird kind: ", err.Error())
	}
	objsByts, _ = json.Marshal(objs)
	fmt.Println("Birds: ", string(objsByts))
}

```
