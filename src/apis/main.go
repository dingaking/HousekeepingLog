package main

import (
	"apis/query"
	"log"
	"net/http"
)

func main() {

	initAdminUserIfNeeded()
	initSystemConfiguration()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8082", router))
}

// admin/admin 사용자가 없는 경우 추가한다.
func initAdminUserIfNeeded() {
	session, err := query.GetConnect()
	if err != nil {
		panic(err)
	}
	query.AdminInitOnBoot(session, "user")
}

// 시스템 설정
func initSystemConfiguration() {
	session, err := query.GetConnect()
	if err != nil {
		panic(err)
	}
	query.AdminSystemConfigurationOnBoot(session, "admin")
}
