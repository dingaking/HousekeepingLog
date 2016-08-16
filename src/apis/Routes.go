package main

import (
	"apis/handler"
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
		"AuthC",
		"POST",
		"/api/authC",
		handler.AuthC,
	},
	Route{
		"AuthR",
		"POST",
		"/api/authR",
		handler.AuthR,
	},
	Route{
		"CapitalC",
		"POST",
		"/api/capitalC",
		handler.CapitalC,
	},
	Route{
		"CategoryC",
		"POST",
		"/api/categoryC",
		handler.CategoryC,
	},
	Route{
		"DeviceC",
		"POST",
		"/api/deviceC",
		handler.DeviceC,
	},
	Route{
		"ItemC",
		"POST",
		"/api/itemC",
		handler.ItemC,
	},
	Route{
		"SystemC",
		"POST",
		"/api/systemC",
		handler.SystemC,
	},
}
