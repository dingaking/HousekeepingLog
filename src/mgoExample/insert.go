package main

import (
	"time"

	"labix.org/v2/mgo"
)

func sample_insert() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Collection People
	c := session.DB("test").C("people")

	// Insert Datas
	err = c.Insert(&Person{Name: "Ale", Phone: "+55 53 1234 4321", Timestamp: time.Now()},
		&Person{Name: "Cla", Phone: "+66 33 1234 5678", Timestamp: time.Now()})

	if err != nil {
		panic(err)
	}
}
