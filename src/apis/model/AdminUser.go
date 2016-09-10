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
