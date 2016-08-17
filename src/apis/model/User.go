package model

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type User struct {
	UserNo         bson.ObjectId `bson:"_id,omitempty"`
	UserId         string        `json:"userid"`
	Password       string        `json:"password"`
	UserType       int           `json:"usertype"`
	DisplayName    string        `json:"displayname"`
	Intro          string        `json:"intro"`
	Profile        string        `json:"profile"`
	CreateDateTime time.Time     `json"create_datetime"`
	LastAccess     time.Time     `json:"last_access"`
	PhoneNumber    string        `json:"phone_number"`
	State          int           `json:"state"`
	Activated      int           `json:"activated"`
	GroupnoList    []int         `json:"groupno_list"`
	TokenList      []TokenInfo   `json:"tokeninfo_list"`
}

type TokenInfo struct {
	TerminalName   string    `json:"terminal_name"`
	TerminalId     string    `json:"terminal_id"`
	TokenId        string    `json:"token_id"`
	ExpireDateTime time.Time `json:"expire_datetime"`
}
