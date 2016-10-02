package model

type AdminUserRReq struct {
	Action      string `json:"action"`
	UserNo      string `json:"userno"`
	UserId      string `json:"userid"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

type AdminUserRRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
	AccessToken  string `json:"access_token"`
	Data         UserJ  `json:"data"`
}

type AdminUserRRes1 struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
	AccessToken  string `json:"access_token"`
}

type AdminUserRRes2 struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
	Data         UserJ  `json:"data"`
}

type AdminUserUReq struct {
	Action      string     `json:action`
	UserId      string     `json:"userid"`
	OldPassword string     `json:"old_password"`
	NewPassword string     `json:"new_password"`
	UserNo      string 	   `json:"userno"`
	AccessToken string     `json:"access_token"`
	UserUpdate  UserUpdate `json:"update"`
}

type AdminUserURes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
	AccessToken  string `json:"access_token"`
}

type AdminUserCReq struct {
	UserId      string `json:"userid"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

type AdminUserCRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type AdminUserLReq struct {
	AccessToken string `json:"access_token"`
}

type AdminUserLRes struct {
	Result       string  `json:"result"`
	ErrorMessage string  `json:"err_msg"`
	Data         []UserJ `json:"data"`
}

type AdminUserDReq struct {
	AccessToken string `json:"access_token"`
	UserNo      string `json:"userno"`
}

type AdminUserDRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type AdminUserSReq struct {
	AccessToken string     `json:"access_token"`
	Page        string     `json:"page"`
	Search      UserSearch `json:"search"`
}

type AdminUserSRes struct {
	Result       string  `json:"result"`
	ErrorMessage string  `json:"err_msg"`
	Data         []UserJ `json:"data"`
}
