package checker

import (
	"apis/model"
	"errors"
)

func AdminUserR(req model.AdminUserRReq) error {

	if len(req.UserId) <= 0 {
		return errors.New("UserId not found.")
	}
	if len(req.Password) <= 0 {
		return errors.New("Password not found.")
	}

	return nil
}
