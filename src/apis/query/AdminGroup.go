package query

import (
	"apis/model"
	"errors"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func CheckPermission(s *mgo.Session, db string, collection string, req *model.AdminGroupCReq) error {
	c := s.DB(db).C(collection)

	var result model.User
	c.Find(bson.M{"access_token": req.AccessToken}).One(&result)

	if result.UserNo == "" {
		return errors.New("wrong access token error.")
	}
	if result.UserType != 4 {
		return errors.New("not admin error.")
	}

	return nil
}

func InsertGroup(s *mgo.Session, db string, collection string, req *model.AdminGroupCReq) error {
	c := s.DB(db).C(collection)

	var insert = model.Group{
		GroupName:      req.GroupName,
		CreateDateTime: time.Now(),
		State:          1,
	}

	err := c.Insert(insert)
	return err
}
