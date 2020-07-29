package controllers

import (
	"fmt"
	model "gitchat/models"
	util "gitchat/utils"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginGet (c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})

}

//登录
func LoginPost(c *gin.Context) {
	//获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username, ",password:", password)

	id := model.QueryUserWithParam(username, util.MD5(password))
	fmt.Println("id:", id)
	if id > 0 {

		session := sessions.Default(c)
		session.Set("loginuser", username)
		session.Save()

		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})
		c.Redirect(http.StatusMovedPermanently,"/")

	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}
}