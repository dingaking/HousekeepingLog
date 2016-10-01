package query

import (
	"apis/model"
	"apis/util"
	"errors"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func AuthC(s *mgo.Session, collection string, authc *model.AuthCReq) error {

	err := checkDuplicationUserId(s, collection, authc)
	if err != nil {
		return err
	}

	err = insertUserId(s, collection, authc)
	return err
}

// insert user document
// 1. make access token and set value to AuthCReq.AccessToken field
// 2. insert info to DB
func insertUserId(s *mgo.Session, collection string, authc *model.AuthCReq) error {

	c := s.DB(DatabaseName).C(collection)

	authc.AccessToken = util.SHA1()

	var insert = model.User{
		UserId:         authc.UserId,
		Password:       authc.Password,
		UserType:       5,
		CreateDateTime: time.Now(),
		State:          1,
		Activated:      0,
		Public:         0,
		AccessToken:    authc.AccessToken,
	}

	err := c.Insert(&insert)
	return err
}

// userid duplication check
// return nil : success
// return error : duplicated
func checkDuplicationUserId(s *mgo.Session, collection string, authc *model.AuthCReq) error {

	c := s.DB(DatabaseName).C(collection)

	var result model.User
	c.Find(bson.M{"userid": authc.UserId}).One(&result)
	if result.UserNo != "" {
		return errors.New("duplicated userid error.")
	}

	return nil
}

func GetUserInfoById(s *mgo.Session, collection string, req *model.AuthRReq) model.User {
	var result model.User
	c := s.DB(DatabaseName).C(collection)
	c.Find(bson.M{"userid": req.UserId, "password": req.Password}).One(&result)
	return result
}

func GetUserInfoByAccessToken(s *mgo.Session, collection string, req *model.AuthRReq) model.User {
	var result model.User
	c := s.DB(DatabaseName).C(collection)
	c.Find(bson.M{"accesstoken": req.AccessToken}).One(&result)
	return result
}

func AuthU(s *mgo.Session, req model.AuthUReq, rep *model.AuthURes) error {

	return nil
}

func AuthD(s *mgo.Session, req model.AuthDReq, rep *model.AuthDRes) error {

	return nil
}
