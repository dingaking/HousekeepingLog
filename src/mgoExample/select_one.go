package main

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func sample_select_one() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Collection People
	c := session.DB("test").C("people")

	// Query One
	result := Person{}
	err = c.Find(bson.M{"name": "Ale"}).Select(bson.M{"phone": 0}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone", result)
}
