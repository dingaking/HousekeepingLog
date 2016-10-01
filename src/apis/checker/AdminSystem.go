package checker

import (
	"apis/model"
	"errors"
)

func AdminSystemL(req model.AdminSystemLReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func AdminSystemR(req model.AdminSystemRReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	if len(req.AdminNo) <= 0 {
		return errors.New("AdminNo not found.")
	}
	return nil
}

func AdminSystemS(req model.AdminSystemSReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	if len(req.Search) <= 0 {
		return errors.New("search valuenot found.")
	}
	return nil
}

func AdminSystemU(req model.AdminSystemUReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	if len(req.AdminNo) <= 0 {
		return errors.New("AdminNo not found.")
	}
	if len(req.ItemValue) <= 0 {
		return errors.New("ItemValue not found.")
	}

	return nil
}
