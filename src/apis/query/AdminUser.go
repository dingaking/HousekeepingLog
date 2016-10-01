package query

import (
	"apis/model"
	"apis/util"
	"errors"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func AdminUserR(s *mgo.Session, req *model.AdminUserRReq) (error, string) {

	c := s.DB(DatabaseName).C(CollUser)

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
func AdminUserU(s *mgo.Session, req *model.AdminUserUReq) error {
	c := s.DB(DatabaseName).C(CollUser)

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

func AdminUserC(s *mgo.Session, req model.AdminUserCReq, rep *model.AdminUserCRes) error {

	err := checkDuplicationUserIdFromAdmin(s, req)
	if err != nil {
		return err
	}

	err = insertUserIdFromAdmin(s, req)
	return err
}

func AdminUserL(s *mgo.Session, req model.AdminUserLReq, rep *model.AdminUserLRes) error {

	c := s.DB(DatabaseName).C(CollUser)

	var data []model.User
	err := c.Find(bson.M{}).All(&data)

	rep.Data = make([]model.UserJ, 0, 0)
	for _, user := range data {
		rep.Data = append(rep.Data, model.UserJ{
			UserNo:         user.UserNo.Hex(),
			UserId:         user.UserId,
			UserType:       user.UserType,
			DisplayName:    user.DisplayName,
			Intro:          user.Intro,
			Profile:        user.Profile,
			CreateDateTime: user.CreateDateTime,
			PhoneNumber:    user.PhoneNumber,
			State:          user.State,
			Activated:      user.Activated,
			Public:         user.Public,
		})
	}

	return err
}

func AdminUserD(s *mgo.Session, req model.AdminUserDReq, rep *model.AdminUserDRes) error {

	return nil
}

func AdminUserS(s *mgo.Session, req model.AdminUserSReq, rep *model.AdminUserSRes) error {

	return nil
}

// userid duplication check
// return nil : success
// return error : duplicated
func checkDuplicationUserIdFromAdmin(s *mgo.Session, req model.AdminUserCReq) error {

	c := s.DB(DatabaseName).C(CollUser)

	var result model.User
	c.Find(bson.M{"userid": req.UserId}).One(&result)
	if result.UserNo != "" {
		return errors.New("duplicated userid error.")
	}

	return nil
}

// insert user document
// 1. make access token and set value to AccessToken field
// 2. insert info to DB
func insertUserIdFromAdmin(s *mgo.Session, req model.AdminUserCReq) error {

	c := s.DB(DatabaseName).C(CollUser)

	req.AccessToken = util.SHA1()

	var insert = model.User{
		UserId:         req.UserId,
		Password:       req.Password,
		UserType:       5,
		CreateDateTime: time.Now(),
		State:          1,
		Activated:      0,
		Public:         0,
		AccessToken:    req.AccessToken,
	}

	err := c.Insert(&insert)
	return err
}
