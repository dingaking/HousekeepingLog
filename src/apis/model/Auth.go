/*
model package : 가계부에서 사용하는 모델에 대한 정의를 담고 있습니다.
*/
package model

import (
	"time"

	"labix.org/v2/mgo/bson"
)

// 회원 모델
type User struct {
	UserNo         bson.ObjectId `bson:"_id,omitempty"`   // 회원키
	UserId         string        `bson:"userid"`          // userid(email)
	Password       string        `bson:"password"`        // 비밀번호(8~16자리)
	UserType       int           `bson:"usertype"`        // 5:사용자, 9:관리자
	DisplayName    string        `bson:"displayname"`     // 대화명
	Intro          string        `bson:"intro"`           // 자기소개
	Profile        string        `bson:"profile"`         // 프로필 사진
	CreateDateTime time.Time     `bson:"create_datetime"` // 가입일시
	PhoneNumber    string        `bson:"phone_number"`    // 전화번호
	State          int           `bson:"state"`           // 회원상태, 1=ON(사용), 0=OFF(탈퇴)
	Activated      int           `bson:"activated"`       // 인증여부, 2=관리자인증, 1=URL인증, 0=OFF(미인증)
	Public         int           `bson:"public"`          // 그룹에 공개여부, 1=공개 0=비공개
	AccessToken    string        `bson:"access_token"`    // 회원접속키
}

type UserJ struct {
	UserNo         string    `json:"userno"`          // 회원키
	UserId         string    `json:"userid"`          // userid(email)
	UserType       int       `json:"usertype"`        // 5:사용자, 9:관리자
	DisplayName    string    `json:"displayname"`     // 대화명
	Intro          string    `json:"intro"`           // 자기소개
	Profile        string    `json:"profile"`         // 프로필 사진
	CreateDateTime time.Time `json:"create_datetime"` // 가입일시
	PhoneNumber    string    `json:"phone_number"`    // 전화번호
	State          int       `json:"state"`           // 회원상태, 1=ON(사용), 0=OFF(탈퇴)
	Activated      int       `json:"activated"`       // 인증여부, 2=관리자인증, 1=URL인증, 0=OFF(미인증)
	Public         int       `json:"public"`          // 그룹에 공개여부, 1=공개 0=비공개
	AccessToken    string    `json:"access_token"`    // 회원접속키
}

type UserSearch struct {
	UserId      string `json:"userid"`      // userid(email)
	UserType    int    `json:"usertype"`    // 5:사용자, 9:관리자
	DisplayName string `json:"displayname"` // 대화명
	Intro       string `json:"intro"`       // 자기소개
	Profile     string `json:"profile"`     // 프로필 사진
	date_begin  string `json:"date_begin"`
	date_end    string `json:"date_end"`
	PhoneNumber string `json:"phone_number"` // 전화번호
	State       int    `json:"state"`        // 회원상태, 1=ON(사용), 0=OFF(탈퇴)
	Activated   int    `json:"activated"`    // 인증여부, 2=관리자인증, 1=URL인증, 0=OFF(미인증)
	Public      int    `json:"public"`       // 그룹에 공개여부, 1=공개 0=비공개
}

type UserUpdate struct {
	UserType    string `json:"usertype"`     // 5:사용자, 4:관리자
	DisplayName string `json:"displayname"`  // 대화명
	Intro       string `json:"intro"`        // 자기소개
	PhoneNumber string `json:"phone_number"` // 전화번호
	State       string `json:"state"`        // 회원상태, 1=ON(사용), 0=OFF(탈퇴)
	Activated   string `json:"activated"`    // 인증여부, 2=관리자인증, 1=URL인증, 0=OFF(미인증)
	Public      string `json:"public"`       // 그룹에 공개여부, 1=공개 0=비공개
}

// 단말기 모델
type Terminal struct {
	TerminalNo     bson.ObjectId `bson:"_id,omitempty"`   // 단말키
	UserNo         string        `bson:"userno"`          // 회원키
	TokenId        string        `bson:"token_id"`        // 접속인증키(서버에서 생성)
	TerminalId     string        `bson:"terminal_id"`     // 터미널ID(단말에서 생성,Web은 null)
	DeviceType     int           `bson:"device_type"`     // 단말기 종류, i:iOS, a:android, w:web
	DeviceToken    string        `bson:"device_token"`    // type이 i,a인 경우의 PUSH 키
	CreateDateTime time.Time     `bson:"create_datetime"` // 단말기 등록일시
	Notification   int           `bson:"notification"`    // PUSH 수신 여부, 1:ON, 0:OFF
	State          int           `bson:"state"`           // 사용여부, 1=ON(사용), 0=OFF(삭제)
	TerminalName   string        `bson:"terminal_name"`   // 단말기 별칭
	ExpireDateTime time.Time     `bson:"expire_datetime"` // token_id 만료일시
	LastAccess     time.Time     `bson:"last_access"`     // 전화번호
}

// 관리자 모델 <---->
// 시스템에서 사용하는 관리 항목에 대한 정보를 처리합니다.
// email_addr, email_pass, email_smtp, email_port 등등.
type Admin struct {
	SystemNo  bson.ObjectId `bson:"_id,omitempty"` // 관리 항목키
	ItemKey   string        `json:"item_key"`      // 관리항목
	ItemValue string        `json:"item_value"`    // 관리내용
	State     int           `json:"state"`         // 사용여부, 1=ON(사용), 0=OFF(삭제)
}

// 회원가입 요청 Model
type AuthCReq struct {
	UserId      string `json:"userid"`
	Password    string `json:"password"`
	TerminalId  string `json:"terminal_id"`
	DeviceType  int    `json:"device_type"`
	AccessToken string `json:"access_token"`
}

type AuthCRep struct {
	ErrorMessage string `json:"err_msg"`
	Result       string `json:"result"`
	AccessToken  string `json:"access_token"`
}

type AuthRReq struct {
	Action      string `json:"action"`
	UserId      string `json:"userid"`
	Password    string `json:"password"`
	UserNo      string `json:"userno"`
	AccessToken string `json:"access_token"`
}

type AuthRRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
	AccessToken  string `json:"access_token"`
}

type AuthUReq struct {
	AccessToken string `json:"access_token"`
}

type AuthURes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type AuthDReq struct {
	AccessToken string `json:"access_token"`
}

type AuthDRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
