package checker

import (
	"apis/model"
	"errors"
)

func SystemR(req model.SystemRReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}
