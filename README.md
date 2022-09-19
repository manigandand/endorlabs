## Store Layer:

Implements the `ObjectDB` store interface. This implements adapter pattern.

> Directory struct

- api
  -- containts the basic api interface to test the `ObjectDB` store interface.
- config
  -- app configs
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

Please refer the API documentations:

https://documenter.getpostman.com/view/1310922/2s7YtXht6d

Postman collection:
https://www.getpostman.com/collections/c528dc355be31338e34a

> Known Issue:

kind will be the type of go struct. Ex: `*schema.Person`

1. Save person:

```sh
curl --location --request POST 'http://localhost:8080/v1/objects/persons' \
--data-raw '{
    "name": "Manigandan",
    "last_name": "Dharmalingam",
    "birth_day": "22/07/1993",
    "birth_date": "1993-07-22T05:04:05.000Z"
}'

Res: 201
{
    "name": "Manigandan",
    "id": "2d6b911e-2ad6-4213-8499-e30e52034bbf",
    "last_name": "Dharmalingam",
    "birth_day": "22/07/1993",
    "birth_date": "1993-07-22T05:04:05Z"
}
```

2. Save Animal:

```sh
curl --location --request POST 'http://localhost:8080/v1/objects/animals' \
--data-raw '{
    "name": "scooby",
    "type": "dog",
    "owner_id": "manigandan"
}'

Res: 201
{
    "name": "scooby",
    "id": "246773bb-94cf-4190-99fe-641fcd85f236",
    "type": "dog",
    "owner_id": "manigandan"
}
```

3. Get object by id

```sh
curl --location --request GET 'http://localhost:8080/v1/objects/246773bb-94cf-4190-99fe-641fcd85f236'

Res: 200
{
    "name": "scooby",
    "id": "246773bb-94cf-4190-99fe-641fcd85f236",
    "type": "dog",
    "owner_id": "manigandan"
}
```

4. Delete object

```sh
curl --location --request DELETE 'http://localhost:8080/v1/objects/1a4e662b-b347-4537-9b1d-27a08f8d6368'

Res: 204
```

5. Get object by name

```sh
curl --location --request GET 'http://localhost:8080/v1/objects/name/Manigandan'

Res: 200
{
    "name": "Manigandan",
    "id": "b715b476-73fb-4bb6-962d-0c4459dc090f",
    "last_name": "Dharmalingam",
    "birth_day": "22/07/1993",
    "birth_date": "1993-07-22T05:04:05Z"
}
```

6. List objects by kind

kind: oneof[`*schema.Person`, `*schema.Animal`]

```sh
curl --location --request GET 'http://localhost:8080/v1/objects/kind/*schema.Person'

Res: 200
[
    {
        "name": "Manigandan",
        "id": "b715b476-73fb-4bb6-962d-0c4459dc090f",
        "last_name": "Dharmalingam",
        "birth_day": "22/07/1993",
        "birth_date": "1993-07-22T05:04:05Z"
    }
]
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
