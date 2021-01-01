package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_wallpaper/api"
	"go_wallpaper/middleware"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		fmt.Println("getting reques ...t")
		v1.GET("papular", api.Fetch500pxPapular)

		// 需要登录保护的
		//auth := v1.Group("")
		//auth.Use(middleware.AuthRequired())
		//{
		//	// User Routing
		//	auth.GET("user/me", api.UserMe)
		//	auth.DELETE("user/logout", api.UserLogout)
		//}
	}
	return r
}
