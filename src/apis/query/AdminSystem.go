package query

import (
	"apis/model"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func SystemL(s *mgo.Session, db string, collection string, rep *model.AdminSystemLRep) error {
	c := s.DB(db).C(collection)

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
