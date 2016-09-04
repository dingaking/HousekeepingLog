package checker

import (
	"apis/model"
	"errors"
	"fmt"
	"regexp"
)

func AuthC(user model.AuthCReq) error {

	if len(user.UserId) <= 0 {
		return errors.New("UserId not found.")
	}

	if !validateEmail(user.UserId) {
		fmt.Println("adfadfadfadsf")
		return errors.New("invalid email address.")
	}

	if len(user.Password) <= 0 {
		return errors.New("Password not found.")
	}

	if !validatePassword(user.Password) {
		return errors.New("invalid password.")
	}

	return nil
}

func AuthR(request model.User) error {

	if len(request.UserId) <= 0 {
		return errors.New("UserId not found.")
	}

	if len(request.Password) <= 0 {
		return errors.New("Password not found.")
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
