package routers

import (
	//"../controllers"
	_ "GitChat/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	router.GET("/register",controllers.RegisterGet)

	return router;

}
