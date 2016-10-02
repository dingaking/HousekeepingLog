package checker

import (
	"apis/model"
	"errors"
)

func AdminUserR(req model.AdminUserRReq) error {

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
			return errors.New("AccessToken not found.")
		}
		if len(req.UserNo) <= 0 {
			return errors.New("Password not found.")
		}
	}

	return nil
}

// action=1(비밀번호 변경)
// action=2(회원정보 변경)
// action=3(프로필 이미지 변경)
func AdminUserU(req model.AdminUserUReq) error {

	if len(req.Action) <= 0 {
		return errors.New("Action not found.")
	}
	if req.Action != "1" && req.Action != "2" && req.Action != "3" {
		return errors.New("Action value is Invalid.")
	}

	if req.Action == "1" {
		if len(req.UserId) <= 0 {
			return errors.New("UserId not found.")
		}
		if len(req.OldPassword) <= 0 {
			return errors.New("Password not found.")
		}
		if len(req.NewPassword) <= 0 {
			return errors.New("Password not found.")
		}
	}

	if req.Action == "2" {
		if len(req.AccessToken) <= 0 {
			return errors.New("AccessToken not found.")
		}
		if len(req.UserNo) <= 0 {
			return errors.New("Password not found.")
		}

		if len(req.UserUpdate.UserType) <= 0 &&
			len(req.UserUpdate.DisplayName) <= 0 &&
			len(req.UserUpdate.Intro) <= 0 &&
			len(req.UserUpdate.PhoneNumber) <= 0 &&
			len(req.UserUpdate.State) <= 0 &&
			len(req.UserUpdate.Activated) <= 0 &&
			len(req.UserUpdate.Public) <= 0 {
			return errors.New("update value not found.")
		}
		
		if len(req.UserUpdate.UserType) > 0 {
			if req.UserUpdate.UserType != "4" && req.UserUpdate.UserType != "5" {
				return errors.New("UserType value is Invalid.")
			}
		}
		if len(req.UserUpdate.State) > 0 {
			if req.UserUpdate.State != "0" && req.UserUpdate.State != "1" {
				return errors.New("State value is Invalid.")
			}
		}
		if len(req.UserUpdate.Activated) > 0 {
			if req.UserUpdate.Activated != "1" && req.UserUpdate.Activated != "2" && req.UserUpdate.Activated != "3" {
				return errors.New("Activated value is Invalid.")
			}
		}
		if len(req.UserUpdate.Public) > 0 {
			if req.UserUpdate.Public != "0" && req.UserUpdate.Public != "1" {
				return errors.New("Public value is Invalid.")
			}
		}
	}

	if req.Action == "3" {
		if len(req.AccessToken) <= 0 {
			return errors.New("AccessToken not found.")
		}
		if len(req.UserNo) <= 0 {
			return errors.New("Password not found.")
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
	if len(req.UserNo) <= 0 {
		return errors.New("UserNo not found.")
	}

	return nil
}

func AdminUserS(req model.AdminUserSReq) error {

	if len(req.AccessToken) <= 0 {
		return errors.New("AccessToken not found.")
	}

	if len(req.Search.UserId) <= 0 {
		return errors.New("search data not found.")
	}

	return nil
}
