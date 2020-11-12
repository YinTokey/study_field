package routes

import (
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/v12/_examples/structuring/bootstrap/bootstrap"
	"lottery/web/controllers"
)

func Configure(b *bootstrap.Bootstrapper) {

	index := mvc.New(b.Application.Party())
	index.Handle(new(controllers.IndexController))

}