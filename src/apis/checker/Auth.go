package checker

import "apis/model"

func AuthC(request model.User) string {

	if len(request.UserId) <= 0 {
		return "UserId not found."
	}

	if len(request.Password) <= 0 {
		return "Password not found."
	}

	return ""
}

func AuthR(request model.User) string {

	if len(request.UserId) <= 0 {
		return "UserId not found."
	}

	if len(request.Password) <= 0 {
		return "Password not found."
	}

	return ""
}
