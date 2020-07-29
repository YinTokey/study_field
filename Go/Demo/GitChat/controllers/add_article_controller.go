package controllers

import (
	"fmt"
	model "gitchat/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AddArticleGet(c *gin.Context) {

	//获取session
	islogin := GetSession(c)

	c.HTML(http.StatusOK, "write_article.html", gin.H{"IsLogin": islogin})
}

func AddArticlePost(c *gin.Context) {

	//获取浏览器传输的数据，通过表单的name属性获取值
	//获取表单信息
	title := c.PostForm("title")
	tags := c.PostForm("tags")
	short := c.PostForm("short")
	content := c.PostForm("content")
	fmt.Printf("title:%s,tags:%s\n", title, tags)

	//实例化model，将它出入到数据库中

	art := model.Article{0, title, tags, short, content, "yin", time.Now().Unix()}
	_, err := model.AddArticle(art)

	//返回数据给浏览器
	response := gin.H{}
	if err == nil {
		//无误
		response = gin.H{"code": 1, "message": "ok"}
	} else {
		response = gin.H{"code": 0, "message": "error"}
	}

	c.JSON(http.StatusOK, response)

}
