package query

import (
	"apis/model"

	"labix.org/v2/mgo"
)

func TerminalC(session *mgo.Session, db string, collection string, terminal model.Terminal) error {

	err := insertTerminal(session, db, collection, terminal)
	return err
}

func insertTerminal(session *mgo.Session, db string, collection string, terminal model.Terminal) error {

	c := session.DB(db).C(collection)

	var insert = model.Terminal{}

	err := c.Insert(&insert)
	return err
}
