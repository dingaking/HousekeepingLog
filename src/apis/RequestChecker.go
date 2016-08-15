package main

import "apis/model"

func CheckAuthC(request model.User) string {

	if len(request.UserId) <= 0 {
		return "UserId not found."
	}

	if len(request.Password) <= 0 {
		return "Password not found."
	}

	return ""
}

func CheckAuthR(request model.User) string {

	if len(request.UserId) <= 0 {
		return "UserId not found."
	}

	if len(request.Password) <= 0 {
		return "Password not found."
	}

	return ""
}
