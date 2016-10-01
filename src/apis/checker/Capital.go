package checker

import (
	"apis/model"
	"errors"
)

func CapitalC(request model.CapitalCReq) error {

	if len(request.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}

func CapitalR(req model.CapitalRReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}

func CapitalU(req model.CapitalUReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}

func CapitalD(req model.CapitalDReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}

func CapitalL(req model.CapitalLReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}

func CapitalS(req model.CapitalSReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}
