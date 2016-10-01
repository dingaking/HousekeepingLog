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

func AdminUserU(req model.AdminUserUReq) error {
	if len(req.Action) <= 0 {
		return errors.New("action not found.")
	}
	if req.Action == "1" {
		if len(req.UserId) <= 0 {
			return errors.New("userid not found.")
		}
		if len(req.OldPassword) <= 0 {
			return errors.New("previous password not found.")
		}
		if len(req.NewPassword) <= 0 {
			return errors.New("New password not found.")
		}
	}
	return nil
}

func AdminUserC(req model.AdminUserCReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	if len(req.UserId) <= 0 {
		return errors.New("UserId not found.")
	}

	if !validateEmail(req.UserId) {
		return errors.New("invalid userid(email) address.")
	}

	if len(req.Password) <= 0 {
		return errors.New("Password not found.")
	}

	if !validatePassword(req.Password) {
		return errors.New("invalid password.")
	}

	return nil
}

func AdminUserL(req model.AdminUserLReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}

func AdminUserD(req model.AdminUserDReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}

func AdminUserS(req model.AdminUserSReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}
