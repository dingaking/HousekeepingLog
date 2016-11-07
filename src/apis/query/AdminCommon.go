package query

import (
	"apis/model"
	"errors"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// CheckPermission : check token is admin.
func CheckPermission(s *mgo.Session, accessToken string) error {

	c := s.DB(DatabaseName).C(CollUser)

	var result model.User
	c.Find(bson.M{"access_token": accessToken}).One(&result)

	if result.UserNo == "" {
		return errors.New("wrong access token error")
	}
	if result.UserType != 4 {
		return errors.New("not admin error")
	}

	return nil
}
