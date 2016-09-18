package model

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type AdminGroupCReq struct {
	AccessToken string `json:"access_token"`
	GroupName   string `json:"group_name"`
}

type Group struct {
	GroupNo        bson.ObjectId `bson:"_id,omitempty"`   // groupno 그룹 no
	GroupName      string        `bson:"group_name"`      // 그룹명
	CreateDateTime time.Time     `bson:"create_datetime"` // 등록일시
	State          int           `bson:"state"`           // 상태, 1=ON(사용), 0=OFF(삭제)
}

type AdminGroupCRep struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
}
