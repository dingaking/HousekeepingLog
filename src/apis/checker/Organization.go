package checker

import (
	"apis/model"
	"errors"
)

func OrganizationR(req model.OrganizationRReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func OrganizationL(req model.OrganizationLReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func OrganizationS(req model.OrganizationSReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}
