package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_wallpaper/service"
	"strconv"
)

// UserRegister 用户注册接口
func Fetch500pxPapular(c *gin.Context) {
	//fmt.Println("500px fetch .....")
	//var service service.PxCollectService = service.NewPxCollectService()
	fmt.Println("unsplash fetch ...")
	service := service.NewUmCollectService()

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	if err := c.ShouldBind(&service); err == nil {
		res, err := service.Papular(page, pageSize)
		if err != nil {
			c.JSON(200, "error happend")
		} else {
			c.JSON(200, res)
		}

	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
