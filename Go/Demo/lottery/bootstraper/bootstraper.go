package bootstraper

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"time"
)

type Configurator func(*Bootstrapper)


type Bootstrapper struct {
	Application *iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time

	Sessions *sessions.Sessions
}

func New(appName, appOwner string, cfgs ... Configurator) *Bootstrapper {
	b := &Bootstrapper{
		AppName: appName,
		AppOwner: appOwner,
		AppSpawnDate: time.Now(),
		Application: iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b

}

func (b *Bootstrapper) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

func (b *Bootstrapper) Bootstrap() *Bootstrapper {

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Application.Run(iris.Addr(addr), cfgs...)
}