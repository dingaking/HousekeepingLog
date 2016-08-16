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
		"CapitalR",
		"POST",
		"/api/capitalR",
		handler.CapitalR,
	},
	Route{
		"CapitalU",
		"POST",
		"/api/capitalU",
		handler.CapitalU,
	},
	Route{
		"CapitalD",
		"POST",
		"/api/capitalD",
		handler.CapitalD,
	},
	Route{
		"CapitalL",
		"POST",
		"/api/capitalL",
		handler.CapitalL,
	},
	Route{
		"CapitalS",
		"POST",
		"/api/capitalS",
		handler.CapitalS,
	},

	Route{
		"CategoryC",
		"POST",
		"/api/categoryC",
		handler.CategoryC,
	},
	Route{
		"CategoryR",
		"POST",
		"/api/categoryR",
		handler.CategoryR,
	},
	Route{
		"CategoryU",
		"POST",
		"/api/categoryU",
		handler.CategoryU,
	},
	Route{
		"CategoryD",
		"POST",
		"/api/categoryD",
		handler.CategoryD,
	},
	Route{
		"CategoryL",
		"POST",
		"/api/categoryL",
		handler.CategoryL,
	},
	Route{
		"CategoryS",
		"POST",
		"/api/categoryS",
		handler.CategoryS,
	},

	Route{
		"ItemC",
		"POST",
		"/api/itemC",
		handler.ItemC,
	},
	Route{
		"ItemR",
		"POST",
		"/api/itemR",
		handler.ItemR,
	},
	Route{
		"ItemU",
		"POST",
		"/api/itemU",
		handler.ItemU,
	},
	Route{
		"ItemD",
		"POST",
		"/api/itemD",
		handler.ItemD,
	},
	Route{
		"ItemL",
		"POST",
		"/api/itemL",
		handler.ItemL,
	},
	Route{
		"ItemS",
		"POST",
		"/api/itemS",
		handler.ItemS,
	},

	Route{
		"SystemR",
		"POST",
		"/api/systemR",
		handler.SystemR,
	},
}
