package model

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type Group struct {
	GroupNo        bson.ObjectId `bson:"_id,omitempty"`   // groupno 그룹 no
	GroupName      string        `bson:"group_name"`      // 그룹명
	CreateDateTime time.Time     `bson:"create_datetime"` // 등록일시
	State          int           `bson:"state"`           // 상태, 1=ON(사용), 0=OFF(삭제)
}

type GroupJ struct {
	GroupNo        string    `json:"groupno"`         // groupno 그룹 no
	GroupName      string    `json:"group_name"`      // 그룹명
	CreateDateTime time.Time `json:"create_datetime"` // 등록일시
	State          int       `json:"state"`           // 상태, 1=ON(사용), 0=OFF(삭제)
}

type AdminGroupCReq struct {
	AccessToken string `json:"access_token"`
	GroupName   string `json:"group_name"`
}

type AdminGroupCRep struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
}

type AdminGroupLReq struct {
	AccessToken string `json:"access_token"`
	GroupName   string `json:"group_name"`
}

type AdminGroupLRep struct {
	ErrorMessage string   `json:"err_msg"`
	Result       string   `json:"result"`
	Data         []GroupJ `json:"data"`
}

type AdminGroupDReq struct {
	AccessToken string `json:"access_token"`
	GroupNo     string `json:"groupno"` // groupno 그룹 no
}

type AdminGroupDRep struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
}

type AdminGroupRReq struct {
	AccessToken string `json:"access_token"`
	GroupNo     string `json:"groupno"` // groupno 그룹 no
}

type AdminGroupRRep struct {
	ErrorMessage string   `json:"err_msg"`
	Result       string   `json:"result"`
	Data         []GroupJ `json:"data"`
}
