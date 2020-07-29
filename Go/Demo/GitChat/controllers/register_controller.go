package controllers

import (
	"fmt"
	"gitchat/models"
	util "gitchat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterGet(c *gin.Context){

	c.HTML(http.StatusOK,"register.html",gin.H{"title":"注册页"})

}

// 处理注册信息
func RegisterPost(c *gin.Context) {
	// 获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	//repassword := c.PostForm("repassword")
	fmt.Println("username:", username, ",password:", password)

	id := model.QueryUserWithUsername(username)

	if id > 0 {
		c.JSON(http.StatusOK,gin.H{"code":0,"message":"用户名已经存在"})
		return
	}

	password = util.MD5(password)
	fmt.Println("md5后：",password)


	user := model.User{0,username,password,0,time.Now().Unix()}
	_,err :=model.InsertUser(user)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"code":0,"message":"注册失败"})
	}else{
		c.JSON(http.StatusOK, gin.H{"code":1,"message":"注册成功"})
		c.Redirect(http.StatusMovedPermanently,"/")

	}
}