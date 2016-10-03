package query

import "labix.org/v2/mgo"

var DatabaseServer = "localhost"

// 가계부 데이터베이스(hlog)
var DatabaseName = "hlog"

// 가계부 파일 데이터베이스(hlog_file)
var DatabaseNameFile = "hlog_file"

// 파일 데이터베이스 프로필
var FileCollProfile = "profile_file"

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
