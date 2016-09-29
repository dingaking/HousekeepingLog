////////////////////////////////////////////////////////////
// author    : Jiang Bian
// reference : https://gist.github.com/border/3489566
////////////////////////////////////////////////////////////

package main

import (
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
	sample_search()
	sample_update()
	sample_delete()
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
