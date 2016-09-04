package query

import "labix.org/v2/mgo"

func GetConnect() (*mgo.Session, error) {

	session, err := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)

	return session, err

}
