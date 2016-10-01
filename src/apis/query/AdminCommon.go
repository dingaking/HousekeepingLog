package query

import (
	"apis/model"
	"errors"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func CheckPermission(s *mgo.Session, access_token string) error {

	c := s.DB(DatabaseName).C(CollUser)

	var result model.User
	c.Find(bson.M{"access_token": access_token}).One(&result)

	if result.UserNo == "" {
		return errors.New("wrong access token error.")
	}
	if result.UserType != 4 {
		return errors.New("not admin error.")
	}

	return nil
}
