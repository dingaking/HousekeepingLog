package checker

import (
	"apis/model"
	"errors"
	"strconv"
)

// AdminGroupC : 그룹생성 요청의 파라미터 체크
func AdminGroupC(req model.AdminGroupCReq) error {
	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	if len(req.GroupName) <= 0 {
		return errors.New("GroupName not found.")
	}
	return nil
}

// AdminGroupL : 그룹목록 요청의 파라미터 체크
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

func AdminGroupU(req model.AdminGroupUReq) error {
	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}
	if len(req.GroupNo) <= 0 {
		return errors.New("Group Name not found.")
	}
	if len(req.State) <= 0 && len(req.GroupName) <= 0 {
		return errors.New("Has no update field.")
	}
	if len(req.State) > 0 {
		i, err := strconv.Atoi(req.State)
		if err != nil {
			return errors.New("state value is invalid.")
		}
		if i != 0 && i != 1 {
			return errors.New("state value is invalid.")
		}
	}
	return nil
}
