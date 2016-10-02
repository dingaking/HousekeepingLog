package checker

import (
	"apis/model"
	"errors"
	"regexp"
)

func AuthC(user model.AuthCReq) error {

	if len(user.UserId) <= 0 {
		return errors.New("UserId not found.")
	}

	if !validateEmail(user.UserId) {
		return errors.New("invalid userid(email) address.")
	}

	if len(user.Password) <= 0 {
		return errors.New("Password not found.")
	}

	if !validatePassword(user.Password) {
		return errors.New("invalid password.")
	}

	return nil
}

func AuthR(req model.AuthRReq) error {

	if len(req.Action) <= 0 {
		return errors.New("Action not found.")
	}
	if req.Action != "1" && req.Action != "2" {
		return errors.New("Action value is Invalid.")
	}

	if req.Action == "1" {
		if len(req.UserId) <= 0 {
			return errors.New("UserId not found.")
		}

		if len(req.Password) <= 0 {
			return errors.New("Password not found.")
		}
	}

	if req.Action == "2" {
		if len(req.AccessToken) <= 0 {
			return errors.New("access_token not found.")
		}
		if len(req.UserNo) <= 0 {
			return errors.New("Password not found.")
		}
	}

	return nil
}

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return Re.MatchString(email)
}

func validatePassword(password string) bool {
	Re := regexp.MustCompile("^[a-zA-Z0-9~!@#$%^&*()-=]*$")
	return Re.MatchString(password)
}

func AuthU(req model.AuthUReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}

func AuthD(req model.AuthDReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	return nil
}
