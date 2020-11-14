package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

)

type AdminUserController struct {
	Ctx            iris.Context

}

func (c *AdminUserController) Get() mvc.Result {

	return mvc.View{
		Name: "admin/user.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "gift",
			//"Datalist": datalist,
			//"Total":    total,
		},
		Layout: "admin/layout.html",
	}

}