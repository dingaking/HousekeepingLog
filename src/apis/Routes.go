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

	// 관리자 API
	Route{"AdminWeb", "GET", "/admin/", handler.StaticHandler},
	Route{"AdminWeb", "GET", "/admin/index.html", handler.StaticHandler},
	Route{"AdminWeb", "GET", "/admin/admin.js", handler.StaticHandler},

	Route{"AdminUserR", "POST", "/api/admin/userR", handler.AdminUserR},
	Route{"AdminUserU", "POST", "/api/admin/userU", handler.AdminUserU},

	Route{"AdminGroupC", "POST", "/api/admin/groupC", handler.AdminGroupC},
	Route{"AdminGroupL", "POST", "/api/admin/groupL", handler.AdminGroupL},
	Route{"AdminGroupD", "POST", "/api/admin/groupD", handler.AdminGroupD},
	Route{"AdminGroupR", "POST", "/api/admin/groupR", handler.AdminGroupR},
	Route{"AdminGroupS", "POST", "/api/admin/groupS", handler.AdminGroupS},
	Route{"AdminGroupU", "POST", "/api/admin/groupU", handler.AdminGroupU},

	Route{"AdminSystemL", "POST", "/api/admin/systemL", handler.AdminSystemL},
	Route{"AdminSystemR", "POST", "/api/admin/systemR", handler.AdminSystemR},
	Route{"AdminSystemS", "POST", "/api/admin/systemS", handler.AdminSystemS},
	Route{"AdminSystemU", "POST", "/api/admin/systemU", handler.AdminSystemU},

	Route{"AdminUserC", "POST", "/api/admin/userC", handler.AdminUserC},
	Route{"AdminUserL", "POST", "/api/admin/userL", handler.AdminUserL},
	Route{"AdminUserD", "POST", "/api/admin/userD", handler.AdminUserD},
	Route{"AdminUserR", "POST", "/api/admin/userR", handler.AdminUserR},
	Route{"AdminUserS", "POST", "/api/admin/userS", handler.AdminUserS},
	Route{"AdminUserU", "POST", "/api/admin/userU", handler.AdminUserU},

	// 사용자 API
	Route{"AuthC", "POST", "/api/authC", handler.AuthC},
	Route{"AuthR", "POST", "/api/authR", handler.AuthR},
	Route{"AuthU", "POST", "/api/authU", handler.AuthU},
	Route{"AuthD", "POST", "/api/authD", handler.AuthD},

	Route{"DeviceR", "POST", "/api/deviceR", handler.DeviceR},
	Route{"DeviceU", "POST", "/api/deviceU", handler.DeviceU},
	Route{"DeviceD", "POST", "/api/deviceD", handler.DeviceD},
	Route{"DeviceL", "POST", "/api/deviceL", handler.DeviceL},
	Route{"DeviceS", "POST", "/api/deviceS", handler.DeviceS},

	Route{"CapitalC", "POST", "/api/capitalC", handler.CapitalC},
	Route{"CapitalR", "POST", "/api/capitalR", handler.CapitalR},
	Route{"CapitalU", "POST", "/api/capitalU", handler.CapitalU},
	Route{"CapitalD", "POST", "/api/capitalD", handler.CapitalD},
	Route{"CapitalL", "POST", "/api/capitalL", handler.CapitalL},
	Route{"CapitalS", "POST", "/api/capitalS", handler.CapitalS},

	Route{"CategoryC", "POST", "/api/categoryC", handler.CategoryC},
	Route{"CategoryR", "POST", "/api/categoryR", handler.CategoryR},
	Route{"CategoryU", "POST", "/api/categoryU", handler.CategoryU},
	Route{"CategoryD", "POST", "/api/categoryD", handler.CategoryD},
	Route{"CategoryL", "POST", "/api/categoryL", handler.CategoryL},
	Route{"CategoryS", "POST", "/api/categoryS", handler.CategoryS},

	Route{"ItemC", "POST", "/api/itemC", handler.ItemC},
	Route{"ItemR", "POST", "/api/itemR", handler.ItemR},
	Route{"ItemU", "POST", "/api/itemU", handler.ItemU},
	Route{"ItemD", "POST", "/api/itemD", handler.ItemD},
	Route{"ItemL", "POST", "/api/itemL", handler.ItemL},
	Route{"ItemS", "POST", "/api/itemS", handler.ItemS},

	Route{"SystemR", "POST", "/api/systemR", handler.SystemR},

	Route{"OrganizationL", "POST", "/api/organizationL", handler.OrganizationL},
	Route{"OrganizationR", "POST", "/api/organizationR", handler.OrganizationR},
	Route{"OrganizationS", "POST", "/api/organizationS", handler.OrganizationS},
}
