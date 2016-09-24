package query

import (
	"apis/model"
	"errors"
	"fmt"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func CheckPermission(s *mgo.Session, db string, collection string, access_token string) error {
	c := s.DB(db).C(collection)

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

func GroupL(s *mgo.Session, db string, collection string, rep *model.AdminGroupLRep) error {
	c := s.DB(db).C(collection)

	var data []model.Group
	err := c.Find(bson.M{}).All(&data)

	rep.Data = make([]model.GroupJ, 0, 0)
	for _, group := range data {
		rep.Data = append(rep.Data, model.GroupJ{
			GroupNo:        group.GroupNo.Hex(),
			GroupName:      group.GroupName,
			CreateDateTime: group.CreateDateTime,
			State:          group.State,
		})
	}
	return err
}

func GroupD(s *mgo.Session, db string, collection string, groupno string) error {

	c := s.DB(db).C(collection)

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(groupno)})

	return err
}

func GroupR(s *mgo.Session, db string, collection string, groupno string) (error, []model.GroupJ) {
	c := s.DB(db).C(collection)

	fmt.Println(groupno)

	var data []model.Group
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(groupno)}).All(&data)

	rep_data := make([]model.GroupJ, 0, 0)
	for _, group := range data {
		rep_data = append(rep_data, model.GroupJ{
			GroupNo:        group.GroupNo.Hex(),
			GroupName:      group.GroupName,
			CreateDateTime: group.CreateDateTime,
			State:          group.State,
		})
	}
	return err, rep_data
}
