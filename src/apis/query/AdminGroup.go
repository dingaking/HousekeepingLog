package query

import (
	"apis/model"
	"errors"
	"strconv"
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

func GroupS(s *mgo.Session, db string, collection string, group_name string) (error, []model.GroupJ) {
	c := s.DB(db).C(collection)

	var data []model.Group
	err := c.Find(bson.M{"group_name": bson.M{"$regex": group_name}}).All(&data)

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

func GroupU(s *mgo.Session, db string, collection string, req model.AdminGroupUReq) error {
	c := s.DB(db).C(collection)

	var result model.Group
	c.Find(bson.M{"_id": bson.ObjectIdHex(req.GroupNo)}).One(&result)
	if result.GroupNo == "" {
		return errors.New("invalid groupno.")
	}

	target := bson.M{"_id": result.GroupNo}
	var change bson.M
	if len(req.State) <= 0 {
		change = bson.M{"$set": bson.M{"group_name": req.GroupName}}
	} else if len(req.GroupName) <= 0 {
		i, _ := strconv.Atoi(req.State)
		change = bson.M{"$set": bson.M{"state": i}}
	} else {
		i, _ := strconv.Atoi(req.State)
		change = bson.M{"$set": bson.M{"group_name": req.GroupName, "state": i}}
	}

	err := c.Update(target, change)
	if err != nil {
		return errors.New("update err")
	}

	return nil
}
