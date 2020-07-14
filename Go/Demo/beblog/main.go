package main

import (
	"github.com/astaxie/beego"
//	_ "bebelog/routers"
	_ "./routers"
)

func main() {
	beego.Run()
}

