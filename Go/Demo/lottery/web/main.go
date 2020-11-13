package main

import (
	"fmt"
	_ "github.com/kataras/iris/v12"
	"lottery/bootstraper"
	"lottery/routes"
	//"github.com/tdewolff/minify/v2"
)

var port = 8080

func newApp() *bootstraper.Bootstrapper {
	app := bootstraper.New("Go抽奖系统", "Yin")
	app.Bootstrap()
	app.Configure(routes.Configure)

	return app
}

func main() {

	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))

}