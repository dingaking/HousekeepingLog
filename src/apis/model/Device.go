package model

type DeviceCReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceCRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceRReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceRRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceUReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceURep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceDReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceDRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceLReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceLRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceSReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceSRep struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
