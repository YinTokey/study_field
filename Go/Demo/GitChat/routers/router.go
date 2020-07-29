package routers

import (
	//"../controllers"
	"gitchat/controllers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*")

	//router.POST("/register",controllers.RegisterPost)

	//router.POST("/register")
	//设置session midddleware
	store := cookie.NewStore([]byte("loginuser"))

	// 参数就是一个处理事件的函数
	router.Use(sessions.Sessions("mysession", store))


	router.GET("/register", controllers.RegisterGet)
	router.POST("/register", controllers.RegisterPost)

	router.GET("/login",controllers.LoginGet)
	router.POST("/login",controllers.LoginPost)

	router.GET("/",controllers.HomeGet)
	router.GET("/exit",controllers.ExitGet)

	router.GET("/article/add",controllers.AddArticleGet)
	router.POST("/article/add",controllers.AddArticlePost)




	return router;

}
