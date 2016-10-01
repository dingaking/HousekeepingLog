package query

import (
	"apis/model"
	"apis/util"
	"errors"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func AdminUserR(s *mgo.Session, db string, collection string, req *model.AdminUserRReq) (error, string) {

	c := s.DB(db).C(collection)

	var result model.User
	c.Find(bson.M{"userid": req.UserId, "password": req.Password}).One(&result)

	if result.UserNo == "" {
		return errors.New("wrong userid or password error."), ""
	}
	if result.UserId == "admin" && result.Password == "admin" {
		return nil, "You must change your initial Password."
	}
	req.AccessToken = result.AccessToken

	return nil, ""
}

//
func AdminUserU(s *mgo.Session, db string, collection string, req *model.AdminUserUReq) error {
	c := s.DB(db).C(collection)

	var result model.User
	c.Find(bson.M{"userid": req.UserId, "password": req.OldPassword}).One(&result)
	if result.UserNo == "" {
		return errors.New("invalid userid or password.")
	}

	req.AccessToken = util.SHA1()

	colQuerier := bson.M{"_id": result.UserNo}
	change := bson.M{"$set": bson.M{"password": req.NewPassword, "accesstoken": req.AccessToken}}
	err := c.Update(colQuerier, change)
	if err != nil {
		return errors.New("update err")
	}

	return nil
}

func AdminUserC(s *mgo.Session, db string, req model.AdminUserCReq, rep *model.AdminUserCRes) error {

	return nil
}

func AdminUserL(s *mgo.Session, db string, req model.AdminUserLReq, rep *model.AdminUserLRes) error {

	return nil
}

func AdminUserD(s *mgo.Session, db string, req model.AdminUserDReq, rep *model.AdminUserDRes) error {

	return nil
}

func AdminUserS(s *mgo.Session, db string, req model.AdminUserSReq, rep *model.AdminUserSRes) error {

	return nil
}
