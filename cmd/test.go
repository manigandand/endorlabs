package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/manigandand/endorlabs/config"
	"github.com/manigandand/endorlabs/schema"
	"github.com/manigandand/endorlabs/store"
)

var (
	dbType = flag.String("store", "inmemory", "store choice")
)

func main() {
	flag.Parse()
	bg := context.Background()

	db := store.Init(config.DBType)
	if db == nil {
		log.Fatalf("🦠store initialize failed 👎")
	}

	per1 := &schema.Person{
		Name:     "Manigandan",
		LastName: "Dharmalingam",
		Birthday: "22/07/1993",
	}

	if err := db.Store(bg, per1); err != nil {
		log.Fatal("cant store person 1: ", err.Error())
	}

	per1Res, err := db.GetObjectByID(bg, per1.ID)
	if err != nil {
		log.Fatal("cant get person 1 by id: ", err.Error())
	}
	fmt.Println(per1Res)

	per1Res, err = db.GetObjectByName(bg, per1.Name)
	if err != nil {
		log.Fatal("cant get person 1 by name: ", err.Error())
	}
	fmt.Println(per1Res)

	ani1 := &schema.Animal{
		Name:    "scooby",
		Type:    "dog",
		OwnerID: "manigandan",
	}
	if err := db.Store(bg, ani1); err != nil {
		log.Fatal("cant store animal 1: ", err.Error())
	}
	ani2 := &schema.Animal{
		Name:    "tucker",
		Type:    "dog",
		OwnerID: "manigandan",
	}
	if err := db.Store(bg, ani2); err != nil {
		log.Fatal("cant store animal 2: ", err.Error())
	}

	objs, err := db.ListObjects(bg, "*schema.Animal")
	if err != nil {
		log.Fatal("cant get objects animals kind: ", err.Error())
	}
	objsByts, _ := json.Marshal(objs)
	fmt.Println("Animals: ", string(objsByts))

	per2 := &schema.Person{
		Name:     "Jeff",
		LastName: "Hardy",
		Birthday: "22/07/1993",
	}

	if err := db.Store(bg, per2); err != nil {
		log.Fatal("cant store person 2: ", err.Error())
	}
	objs, err = db.ListObjects(bg, "*schema.Person")
	if err != nil {
		log.Fatal("cant get objects person kind: ", err.Error())
	}
	objsByts, _ = json.Marshal(objs)
	fmt.Println("Person: ", string(objsByts))
}
