////////////////////////////////////////////////////////////
// author    : Jiang Bian
// reference : https://gist.github.com/border/3489566
////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

var (
	IsDrop = true
)

func main() {

	make_database()
	sample_insert()
	sample_select_one()
	sample_select_all()
	sample_update()
	sample_delete()
}

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

func sample_update() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Collection People
	c := session.DB("test").C("people")

	// Update
	colQuerier := bson.M{"name": "Ale"}
	change := bson.M{"$set": bson.M{"phone": "+86 99 8888 7777", "timestamp": time.Now()}}
	err = c.Update(colQuerier, change)
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

func sample_select_all() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Collection People
	c := session.DB("test").C("people")

	// Query All
	var results []Person
	err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)

	if err != nil {
		panic(err)
	}
	fmt.Println("Results All: ", results)
}

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

func make_database() {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Drop Database
	if IsDrop {
		err = session.DB("test").DropDatabase()
		if err != nil {
			panic(err)
		}
	}

	// Collection People
	c := session.DB("test").C("people")

	// Index
	index := mgo.Index{
		Key:        []string{"name", "phone"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
