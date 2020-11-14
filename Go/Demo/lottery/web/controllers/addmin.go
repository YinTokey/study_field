package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AdminController struct {
	Ctx iris.Context

}

func (c *AdminController) Get() mvc.Result {

	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title": "管理后台",
			"Channel":"",
		},
		Layout: "admin/layout.html",
	}

}