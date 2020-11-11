package main

import (
	"fmt"
	_ "github.com/kataras/iris"
	"lottery/bootstraper"
)

var port = 8080

func newApp() *bootstraper.Bootstrapper {
	app := bootstraper.New("Go抽奖系统", "Yin")
	app.Bootstrap()
	app.Configure()

	return app
}


func main() {

	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))

}