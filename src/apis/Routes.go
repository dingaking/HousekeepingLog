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
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"Index",
		"Get",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"Get",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"Get",
		"/todos/{todoId}",
		TodoShow,
	},
}
