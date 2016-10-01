package model

type DeviceCReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceCRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceRReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceRRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceUReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceURes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceDReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceDRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceLReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceLRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}

type DeviceSReq struct {
	AccessToken string `json:"access_token"`
}

type DeviceSRes struct {
	Result       string `json:"result"`
	ErrorMessage string `json:"err_msg"`
}
