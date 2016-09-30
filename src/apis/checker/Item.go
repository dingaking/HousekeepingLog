package checker

import (
	"apis/model"
	"errors"
)

func ItemC(req model.ItemCReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func ItemR(req model.ItemRReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func ItemU(req model.ItemUReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func ItemD(req model.ItemDReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func ItemL(req model.ItemLReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func ItemS(req model.ItemSReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}
