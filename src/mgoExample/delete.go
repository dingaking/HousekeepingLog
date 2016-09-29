package main

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func sample_delete() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Collection People
	c := session.DB("test").C("people")

	// Delete(Remove)
	err = c.Remove(bson.M{"name": "Ale"})
	if err != nil {
		panic(err)
	}

	// Query All
	var results []Person
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
}
