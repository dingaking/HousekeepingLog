package query

import (
	"apis/model"
	"errors"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func AdminSystemL(s *mgo.Session, rep *model.AdminSystemLRes) error {

	c := s.DB(DatabaseName).C(CollAdmin)

	var data []model.AdminItem
	err := c.Find(bson.M{}).All(&data)

	rep.Data = make([]model.AdminItemJ, 0, 0)
	for _, item := range data {
		rep.Data = append(rep.Data, model.AdminItemJ{
			SystemNo:  item.SystemNo.Hex(),
			ItemKey:   item.ItemKey,
			ItemValue: item.ItemValue,
			ItemDesc:  item.ItemDesc,
		})
	}

	return err
}

func AdminSystemR(s *mgo.Session, adminno string, rep *model.AdminSystemRRes) error {
	c := s.DB(DatabaseName).C(CollAdmin)

	var data model.AdminItem
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(adminno)}).One(&data)

	rep.Data = model.AdminItemJ{
		ItemKey:   data.ItemKey,
		ItemValue: data.ItemValue,
		ItemDesc:  data.ItemDesc,
	}

	return err
}

func AdminSystemS(s *mgo.Session, search string, rep *model.AdminSystemSRes) error {

	c := s.DB(DatabaseName).C(CollAdmin)

	var data []model.AdminItem
	err := c.Find(bson.M{"item_desc": bson.M{"$regex": search}}).All(&data)

	rep.Data = make([]model.AdminItemJ, 0, 0)
	for _, item := range data {
		rep.Data = append(rep.Data, model.AdminItemJ{
			SystemNo:  item.SystemNo.Hex(),
			ItemKey:   item.ItemKey,
			ItemValue: item.ItemValue,
			ItemDesc:  item.ItemDesc,
		})
	}

	return err
}

func AdminSystemU(s *mgo.Session, req model.AdminSystemUReq, rep *model.AdminSystemURes) error {

	c := s.DB(DatabaseName).C(CollAdmin)

	var data model.AdminItem
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(req.SystemNo)}).One(&data)
	if len(data.ItemKey) <= 0 {
		return errors.New("key not found error.")
	}

	target := bson.M{"_id": bson.ObjectIdHex(req.SystemNo)}
	change := bson.M{"$set": bson.M{"item_value": req.ItemValue}}
	if len(req.ItemDesc) > 0 {
		change = bson.M{"$set": bson.M{"item_value": req.ItemValue, "item_desc": req.ItemDesc}}
	}
	err = c.Update(target, change)

	return err
}
