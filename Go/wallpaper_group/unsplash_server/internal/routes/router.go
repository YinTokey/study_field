package routes

import (
	"github.com/gin-gonic/gin"
	"os"
	"unsplash_server/internal/routes/api"
	"unsplash_server/internal/middleware"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	//r.Use(middleware.AccessLog())
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.GET("papular", api.FetchPapular)

	}
	return r
}
