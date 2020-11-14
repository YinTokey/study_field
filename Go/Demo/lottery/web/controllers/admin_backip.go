package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AdminBlackipController struct {
	Ctx iris.Context

}

func (c *AdminBlackipController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/blackip.html",
		Data: iris.Map{
			"Title":    "管理后台",
			"Channel":  "blackip",
			//"Datalist": datalist,
			//"Total":    total,
			//"Now":      comm.NowUnix(),
			//"PagePrev": pagePrev,
			//"PageNext": pageNext,
		},
		Layout: "admin/layout.html",
	}
}