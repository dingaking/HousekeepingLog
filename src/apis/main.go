package main

import (
	"apis/query"
	"log"
	"net/http"
)

func main() {

	initAdminUserIfNeeded()

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8082", router))
}

// admin/admin 사용자가 없는 경우 추가한다.
func initAdminUserIfNeeded() {
	session, err := query.GetConnect()
	if err != nil {
		panic(err)
	}
	query.AdminInitFromBoot(session, "hlog", "user")
}
