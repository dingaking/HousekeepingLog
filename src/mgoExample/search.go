package main

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func sample_search() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Collection People
	c := session.DB("test").C("people")

	// search
	var results []Person
	err = c.Find(bson.M{"name": bson.M{"$regex": "le"}}).Sort("-timestamp").All(&results)
	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
}
