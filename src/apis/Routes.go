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
		"AuthU",
		"POST",
		"/api/authU",
		handler.AuthU,
	},
	Route{
		"AuthD",
		"POST",
		"/api/authD",
		handler.AuthD,
	},

	Route{
		"DeviceC",
		"POST",
		"/api/deviceC",
		handler.DeviceC,
	},
	Route{
		"DeviceR",
		"POST",
		"/api/deviceR",
		handler.DeviceR,
	},
	Route{
		"DeviceU",
		"POST",
		"/api/deviceU",
		handler.DeviceU,
	},
	Route{
		"DeviceD",
		"POST",
		"/api/deviceD",
		handler.DeviceD,
	},
	Route{
		"DeviceL",
		"POST",
		"/api/deviceL",
		handler.DeviceL,
	},
	Route{
		"DeviceS",
		"POST",
		"/api/deviceS",
		handler.DeviceS,
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
