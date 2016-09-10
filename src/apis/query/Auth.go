package query

import (
	"apis/model"
	"apis/util"
	"errors"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func AuthC(session *mgo.Session, db string, collection string, authc *model.AuthCReq) error {

	err := checkDuplicationUserId(session, db, collection, authc)
	if err != nil {
		return err
	}

	err = insertUserId(session, db, collection, authc)
	return err
}

// insert user document
// 1. make access token and set value to AuthCReq.AccessToken field
// 2. insert info to DB
func insertUserId(session *mgo.Session, db string, collection string, authc *model.AuthCReq) error {

	c := session.DB(db).C(collection)

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
func checkDuplicationUserId(session *mgo.Session, db string, collection string, authc *model.AuthCReq) error {

	c := session.DB(db).C(collection)

	var result model.User
	c.Find(bson.M{"userid": authc.UserId}).One(&result)
	if result.UserNo != "" {
		return errors.New("duplicated userid error.")
	}

	return nil
}
