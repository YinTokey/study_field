package server

import (
	"github.com/gin-gonic/gin"
	"Week02/api"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {

	r := gin.Default()

	r.GET("pictures",api.QueryPicture)

	r.GET("add",api.AddPicture)

	return r
}
