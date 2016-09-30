package model

type OrganizationRReq struct {
	AccessToken string `json:"access_token"`
}

type OrganizationRRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type OrganizationLReq struct {
	AccessToken string `json:"access_token"`
}

type OrganizationLRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type OrganizationSReq struct {
	AccessToken string `json:"access_token"`
}

type OrganizationSRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
