package controllers

import (
	"github.com/kataras/iris"
)

type IndexController struct {
	Ctx iris.Context

}

func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go抽奖系统，<a href='/public/index.html'>开始抽奖</a>"
}
