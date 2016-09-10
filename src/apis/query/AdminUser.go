package query

import (
	"apis/model"
	"errors"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func AdminUserR(session *mgo.Session, db string, collection string, req *model.AdminUserRReq) error {

	c := session.DB(db).C(collection)

	var result model.User
	c.Find(bson.M{"userid": req.UserId, "password": req.Password}).One(&result)
	if result.UserNo != "" {
		return errors.New("wrong userid or password error.")
	}

	return nil
}
