package main

import "apis/model"

func CheckAuthC(request model.User) string {

	if len(request.UserId) <= 0 {
		return "UserId not found."
	}

	return ""
}
