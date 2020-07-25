package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterGet(c *gin.Context){

	c.HTML(http.StatusOK,"register.html",gin.H{"title":"注册页"})


}