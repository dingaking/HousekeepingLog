package checker

import (
	"apis/model"
	"errors"
)

func DeviceC(req model.DeviceCReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func DeviceR(req model.DeviceRReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func DeviceU(req model.DeviceUReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func DeviceD(req model.DeviceDReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func DeviceL(req model.DeviceLReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func DeviceS(req model.DeviceSReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}
