package model

type OrganizationRReq struct {
	AccessToken string `json:"access_token"`
}

type OrganizationRRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type OrganizationLReq struct {
	AccessToken string `json:"access_token"`
}

type OrganizationLRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type OrganizationSReq struct {
	AccessToken string `json:"access_token"`
}

type OrganizationSRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
