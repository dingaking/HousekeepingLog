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

func AdminGroupL(req model.AdminGroupLReq) error {
	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	return nil
}

func AdminGroupD(req model.AdminGroupDReq) error {
	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	if len(req.GroupNo) <= 0 {
		return errors.New("Group Number not found.")
	}
	return nil
}

func AdminGroupR(req model.AdminGroupRReq) error {
	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	if len(req.GroupNo) <= 0 {
		return errors.New("Group Number not found.")
	}
	return nil
}

func AdminGroupS(req model.AdminGroupSReq) error {
	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	if len(req.GroupName) <= 0 {
		return errors.New("Group Name not found.")
	}
	return nil
}
