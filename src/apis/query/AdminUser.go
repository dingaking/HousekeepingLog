package query

import (
	"apis/model"
	"apis/util"
	"bufio"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func AdminUserR(s *mgo.Session, req model.AdminUserRReq, rep *model.AdminUserRRes) error {

	c := s.DB(DatabaseName).C(CollUser)

	if req.Action == "1" {
		var data model.User
		err := c.Find(bson.M{"userid": req.UserId, "password": req.Password}).One(&data)
		if data.UserNo == "" {
			return errors.New("user not found error.")
		}
		if data.UserType != 4 {
			return errors.New("Not Admin User error.")
		}
		rep.AccessToken = data.AccessToken

		return err
	} else if req.Action == "2" {
		var data model.User
		err := c.Find(bson.M{"_id": bson.ObjectIdHex(req.UserNo)}).One(&data)

		if data.UserNo == "" {
			return errors.New("user not found error.")
		}

		rep.Data = model.UserJ{
			UserNo:         data.UserNo.Hex(),
			UserId:         data.UserId,
			UserType:       data.UserType,
			DisplayName:    data.DisplayName,
			Intro:          data.Intro,
			Profile:        data.Profile,
			CreateDateTime: data.CreateDateTime,
			PhoneNumber:    data.PhoneNumber,
			State:          data.State,
			Activated:      data.Activated,
			Public:         data.Public,
		}

		return err
	} else {
		return nil
	}
}

// action=1(비밀번호 변경)
// action=2(회원정보 변경)
// action=3(프로필 이미지 변경)
func AdminUserU(s *mgo.Session, req *model.AdminUserUReq) error {

	c := s.DB(DatabaseName).C(CollUser)

	if req.Action == "1" {
		var data model.User
		c.Find(bson.M{"userid": req.UserId, "password": req.OldPassword}).One(&data)
		if data.UserNo == "" {
			return errors.New("invalid userid or password")
		}

		req.AccessToken = util.SHA1()

		target := bson.M{"_id": data.UserNo}
		change := bson.M{"$set": bson.M{"password": req.NewPassword, "access_token": req.AccessToken}}
		err := c.Update(target, change)
		if err != nil {
			return errors.New("update err")
		}
	} else if req.Action == "2" {

		conditions := bson.M{}

		if len(req.UserUpdate.UserType) > 0 {
			conditions["usertype"], _ = strconv.Atoi(req.UserUpdate.UserType)
		}
		if len(req.UserUpdate.DisplayName) > 0 {
			conditions["displayname"] = req.UserUpdate.DisplayName
		}
		if len(req.UserUpdate.Intro) > 0 {
			conditions["intro"] = req.UserUpdate.Intro
		}
		if len(req.UserUpdate.PhoneNumber) > 0 {
			conditions["phone_number"] = req.UserUpdate.PhoneNumber
		}
		if len(req.UserUpdate.State) > 0 {
			conditions["state"], _ = strconv.Atoi(req.UserUpdate.State)
		}
		if len(req.UserUpdate.Activated) > 0 {
			conditions["activated"], _ = strconv.Atoi(req.UserUpdate.Activated)
		}
		if len(req.UserUpdate.Public) > 0 {
			conditions["public"], _ = strconv.Atoi(req.UserUpdate.Public)
		}

		target := bson.M{"_id": bson.ObjectIdHex(req.UserNo)}
		change := bson.M{"$set": conditions}
		err := c.Update(target, change)
		return err
	}

	return nil
}

func AdminUserUAction3(s *mgo.Session, r *http.Request) error {

	db := s.DB(DatabaseNameFile)

	for _, fileHeaders := range r.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			if gridFile, err := db.GridFS(FileCollProfile).Create("filename"); err != nil {
				return err
			} else {
				var userno = r.FormValue("userno")
				gridFile.SetName(fileHeader.Filename)
				gridFile.SetMeta(bson.M{"userno": userno})
				if err := writeToGridFile(file, gridFile); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func writeToGridFile(file multipart.File, gridFile *mgo.GridFile) error {
	reader := bufio.NewReader(file)
	defer func() { file.Close() }()
	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return errors.New("Could not read the input file")
		}
		if n == 0 {
			break
		}
		// write a chunk
		if _, err := gridFile.Write(buf[:n]); err != nil {
			return errors.New("Could not write to GridFs for " + gridFile.Name())
		}
	}
	gridFile.Close()
	return nil
}

func AdminUserC(s *mgo.Session, req model.AdminUserCReq, rep *model.AdminUserCRes) error {

	err := checkDuplicationUserIdFromAdmin(s, req)
	if err != nil {
		return err
	}

	err = insertUserIdFromAdmin(s, req)
	return err
}

func AdminUserL(s *mgo.Session, req model.AdminUserLReq, rep *model.AdminUserLRes) error {

	c := s.DB(DatabaseName).C(CollUser)

	var data []model.User
	err := c.Find(bson.M{}).All(&data)

	rep.Data = make([]model.UserJ, 0, 0)
	for _, user := range data {
		rep.Data = append(rep.Data, model.UserJ{
			UserNo:         user.UserNo.Hex(),
			UserId:         user.UserId,
			UserType:       user.UserType,
			DisplayName:    user.DisplayName,
			Intro:          user.Intro,
			Profile:        user.Profile,
			CreateDateTime: user.CreateDateTime,
			PhoneNumber:    user.PhoneNumber,
			State:          user.State,
			Activated:      user.Activated,
			Public:         user.Public,
		})
	}

	return err
}

func AdminUserD(s *mgo.Session, req model.AdminUserDReq, rep *model.AdminUserDRes) error {

	c := s.DB(DatabaseName).C(CollUser)

	var data model.User
	c.Find(bson.M{"_id": bson.ObjectIdHex(req.UserNo)}).One(&data)

	if data.UserNo == "" {
		return errors.New("user not found error.")
	}

	err := c.Remove(bson.M{"_id": bson.ObjectIdHex(req.UserNo)})
	return err
}

func AdminUserS(s *mgo.Session, req model.AdminUserSReq, rep *model.AdminUserSRes) error {

	c := s.DB(DatabaseName).C(CollUser)

	var data []model.User
	err := c.Find(bson.M{"userid": bson.M{"$regex": req.Search.UserId}}).All(&data)
	if err != nil {
		return err
	}

	rep.Data = make([]model.UserJ, 0, 0)
	for _, user := range data {
		rep.Data = append(rep.Data, model.UserJ{
			UserNo:         user.UserNo.Hex(),
			UserId:         user.UserId,
			UserType:       user.UserType,
			DisplayName:    user.DisplayName,
			Intro:          user.Intro,
			Profile:        user.Profile,
			CreateDateTime: user.CreateDateTime,
			PhoneNumber:    user.PhoneNumber,
			State:          user.State,
			Activated:      user.Activated,
			Public:         user.Public,
		})
	}

	return err
}

// userid duplication check
// return nil : success
// return error : duplicated
func checkDuplicationUserIdFromAdmin(s *mgo.Session, req model.AdminUserCReq) error {

	c := s.DB(DatabaseName).C(CollUser)

	var data model.User
	c.Find(bson.M{"userid": req.UserId}).One(&data)
	if data.UserNo != "" {
		return errors.New("duplicated userid error.")
	}

	return nil
}

// insert user document
// 1. make access token and set value to AccessToken field
// 2. insert info to DB
func insertUserIdFromAdmin(s *mgo.Session, req model.AdminUserCReq) error {

	c := s.DB(DatabaseName).C(CollUser)

	req.AccessToken = util.SHA1()

	var insert = model.User{
		UserId:         req.UserId,
		Password:       req.Password,
		UserType:       5,
		CreateDateTime: time.Now(),
		State:          1,
		Activated:      0,
		Public:         0,
		AccessToken:    req.AccessToken,
	}

	err := c.Insert(&insert)
	return err
}
