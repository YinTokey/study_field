package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_wallpaper/service"
)

// UserRegister 用户注册接口
func Fetch500pxPapular(c *gin.Context) {
	fmt.Println("500px fetch .....")
	var service service.PxCollectService = service.NewPxCollectService()

	if err := c.ShouldBind(&service); err == nil {
		res := service.Papular()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
