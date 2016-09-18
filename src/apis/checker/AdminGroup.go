package checker

import (
	"apis/model"
	"errors"
)

func AdminGroupC(req model.AdminGroupCReq) error {
	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	if len(req.GroupName) <= 0 {
		return errors.New("GroupName not found.")
	}
	return nil
}
