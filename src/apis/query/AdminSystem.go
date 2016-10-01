package query

import (
	"apis/model"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func AdminSystemL(s *mgo.Session, collection string, rep *model.AdminSystemLRep) error {
	c := s.DB(DatabaseName).C(collection)

	var data []model.AdminItem
	err := c.Find(bson.M{}).All(&data)

	rep.Data = make([]model.AdminItemJ, 0, 0)
	for _, item := range data {
		rep.Data = append(rep.Data, model.AdminItemJ{
			AdminNo:   item.AdminNo.Hex(),
			ItemKey:   item.ItemKey,
			ItemValue: item.ItemValue,
			ItemDesc:  item.ItemDesc,
		})
	}

	return err
}

func AdminSystemR(s *mgo.Session, collection string, adminno string, rep *model.AdminSystemRRep) error {
	c := s.DB(DatabaseName).C(collection)

	var data model.AdminItem
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(adminno)}).One(&data)

	rep.Data = model.AdminItemJ{
		ItemKey:   data.ItemKey,
		ItemValue: data.ItemValue,
		ItemDesc:  data.ItemDesc,
	}

	return err
}

func AdminSystemS(s *mgo.Session, collection string, search string, rep *model.AdminSystemSRep) error {
	c := s.DB(DatabaseName).C(collection)

	var data []model.AdminItem
	err := c.Find(bson.M{"item_desc": bson.M{"$regex": search}}).All(&data)

	rep.Data = make([]model.AdminItemJ, 0, 0)
	for _, item := range data {
		rep.Data = append(rep.Data, model.AdminItemJ{
			AdminNo:   item.AdminNo.Hex(),
			ItemKey:   item.ItemKey,
			ItemValue: item.ItemValue,
			ItemDesc:  item.ItemDesc,
		})
	}

	return err
}

func AdminSystemU(s *mgo.Session, req model.AdminSystemUReq, rep *model.AdminSystemSRep) error {

	return nil
}
