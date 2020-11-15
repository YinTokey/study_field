package controllers

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"lottery/comm"
	"lottery/models"
	"lottery/services"
)

type IndexController struct {
	Ctx iris.Context
	ServiceUser services.UserService

}


func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go抽奖系统，<a href='/public/index.html'>开始抽奖</a>"
}

// http://localhost:8080/myprize
//func (c *IndexController) GetMyprize() map[string]interface{} {
//	rs := make(map[string]interface{})
//	rs["code"] = 0
//	rs["msg"] = ""
//	// 验证登录
//	loginuser := comm.GetLoginUser(c.Ctx.Request())
//	if loginuser == nil || loginuser.Uid < 1 {
//		rs["code"] = 101
//		rs["msg"] = "请先登录，再来抽奖"
//		return rs
//	}
//	// 只读取出来最新的100次中奖记录
//	list := c.ServiceResult.SearchByUser(loginuser.Uid, 1, 100)
//	rs["prize_list"] = list
//	// 今天抽奖次数
//	day, _ := strconv.Atoi(comm.FormatFromUnixTimeShort(time.Now().Unix()))
//	num := c.ServiceUserday.Count(loginuser.Uid, day)
//	rs["prize_num"] = conf.UserPrizeMax - num
//	return rs
//}

// 登录 GET /login
func (c *IndexController) GetLogin() {
	// 每次随机生成一个登录用户信息
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/public/index.html?from=login"
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
}

// 退出 GET /logout
func (c *IndexController) GetLogout() {
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/public/index.html?from=logout"
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
}
