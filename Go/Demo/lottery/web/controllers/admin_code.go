package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

)

type AdminCodeController struct {
	Ctx iris.Context

}

func (c *AdminCodeController) Get() mvc.Result {

	return mvc.View{
		Name: "admin/code.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "code",

		},
		Layout: "admin/layout.html",
	}

}