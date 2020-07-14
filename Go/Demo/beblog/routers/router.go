package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	print("aa")

	beego.Router("/", &controllers.MainController{})
}
