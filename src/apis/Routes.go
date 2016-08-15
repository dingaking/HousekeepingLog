package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"UserCreate",
		"POST",
		"/api/authC",
		UserCreate,
	},
	Route{
		"UserInfo",
		"POST",
		"/api/authR",
		UserInfo,
	},
}
