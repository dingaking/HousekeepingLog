package model

type AdminUserRReq struct {
	UserId      string `json:"userid"`
	Password    string `json:"password"`
	AccessToken string `json:"access_token"`
}

type AdminUserRRes struct {
	Result      string `json:"result"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

type AdminUserUReq struct {
	Action      string `json:action`
	UserId      string `json:"userid"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	AccessToken string `json:"access_token"`
}

type AdminUserURes struct {
	Result      string `json:"result"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token"`
}

type AdminUserCReq struct {
	AccessToken string `json:"access_token"`
}

type AdminUserCRes struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

type AdminUserLReq struct {
	AccessToken string `json:"access_token"`
}

type AdminUserLRes struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

type AdminUserDReq struct {
	AccessToken string `json:"access_token"`
}

type AdminUserDRes struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}

type AdminUserSReq struct {
	AccessToken string `json:"access_token"`
}

type AdminUserSRes struct {
	Result  string `json:"result"`
	Message string `json:"message"`
}
