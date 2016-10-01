package query

import "labix.org/v2/mgo"

var DatabaseServer = "localhost"

var CollUser = "user"
var CollTerminal = "terminal"
var CollGroup = "group"
var CollGroupInfo = "group_info"
var CollAdmin = "admin"
var CollCategory = "category"
var CollItem = "item"
var CollCapital = "capital"
var CollUserLog = "user_log"
var CollAdminLog = "admin_log"

func GetConnect() (*mgo.Session, error) {

	session, err := mgo.Dial(DatabaseServer)
	session.SetMode(mgo.Monotonic, true)

	return session, err

}
