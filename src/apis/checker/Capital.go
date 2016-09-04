package checker

import "apis/model"

func CapitalC(request model.Capital) string {

	if len(request.CapitalName) <= 0 {
		return "Capital Name not found."
	}

	return ""
}
