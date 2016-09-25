package query

import (
	"apis/model"
	"errors"
	"fmt"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

// add admin account if has no admin user
func AdminInitOnBoot(s *mgo.Session, db string, collection string) error {

	fmt.Println("/////////////////////////////////////////////////")
	fmt.Println("AdminInitFromBoot")
	fmt.Println("/////////////////////////////////////////////////")

	c := s.DB(db).C(collection)

	var result model.User
	c.Find(bson.M{"userid": "admin"}).One(&result)
	if result.UserNo != "" {
		return nil
	}

	now := time.Now()
	fmt.Println("admin inserted at " + now.String())
	var insert = model.User{
		UserId:         "admin",
		Password:       "admin",
		UserType:       4,
		CreateDateTime: now,
		State:          1,
		Activated:      0,
	}
	err := c.Insert(&insert)
	return err
}

// add system configuration field if needed on boot.
func AdminSystemConfigurationOnBoot(s *mgo.Session, db string, collection string) error {

	c := s.DB(db).C(collection)

	var result []model.AdminItem
	c.Find(bson.M{}).All(&result)

	if len(result) >= 8 {

		var systemTime model.AdminItem
		c.Find(bson.M{"item_key": "system_booting_time"}).One(&systemTime)
		if systemTime.AdminNo == "" {
			return errors.New("system error about boot time.")
		}

		now := time.Now()
		fmt.Println("system booting time inserted at " + now.String())

		target := bson.M{"_id": systemTime.AdminNo}
		var boot = model.AdminItem{
			ItemValue: now.String(),
		}

		err := c.Update(target, boot)
		if err != nil {
			return errors.New("system error about boot time update.")
		}
		return nil
	}

	var item model.AdminItem

	item = model.AdminItem{
		ItemKey:  "email_addr",
		ItemDesc: "smtp sender address.",
	}
	err := c.Insert(item)
	if err != nil {
		return err
	}

	item = model.AdminItem{
		ItemKey:  "email_pass",
		ItemDesc: "smtp sender password.",
	}
	err = c.Insert(item)
	if err != nil {
		return err
	}

	item = model.AdminItem{
		ItemKey:  "email_smtp",
		ItemDesc: "smtp server ip or url.",
	}
	err = c.Insert(item)
	if err != nil {
		return err
	}

	item = model.AdminItem{
		ItemKey:  "email_port",
		ItemDesc: "smtp server port.",
	}
	err = c.Insert(item)
	if err != nil {
		return err
	}

	item = model.AdminItem{
		ItemKey:  "email_success",
		ItemDesc: "smtp url check successed.",
	}
	err = c.Insert(item)
	if err != nil {
		return err
	}

	item = model.AdminItem{
		ItemKey:  "email_authdate",
		ItemDesc: "smtp confirm url sent datetime(string).",
	}
	err = c.Insert(item)
	if err != nil {
		return err
	}

	item = model.AdminItem{
		ItemKey:   "system_version",
		ItemValue: "0.0.1",
		ItemDesc:  "system release version.",
	}
	err = c.Insert(item)
	if err != nil {
		return err
	}

	now := time.Now()
	item = model.AdminItem{
		ItemKey:   "system_booting_time",
		ItemValue: now.String(),
		ItemDesc:  "system booting time.(string)",
	}
	err = c.Insert(item)
	if err != nil {
		return err
	}

	return nil
}
