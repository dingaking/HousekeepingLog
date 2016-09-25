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
