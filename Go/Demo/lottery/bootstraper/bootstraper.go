package bootstraper

import (
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"lottery/conf"
	"time"
)

const (
	// StaticAssets is the root directory for public assets like images, css, js.
	StaticAssets = "./web/public/"
	// Favicon is the relative 9to the "StaticAssets") favicon path for our app.
	Favicon = "favicon.ico"
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

func (b *Bootstrapper) SetupViews(dir string) *Bootstrapper {

	htmlEngine := iris.HTML(dir,"html").Layout("shared/layout.html")

	htmlEngine.Reload(true)

	htmlEngine.AddFunc("FromUnixtimeShort", func(t int) string {
		dt := time.Unix(int64(t),int64(0))
		return dt.Format(conf.SysTimeformShort)
	})
	htmlEngine.AddFunc("FromUnixtime", func(t int) string {
		dt := time.Unix(int64(t),int64(0))
		return dt.Format(conf.SysTimeform)
	})
	b.Application.RegisterView(htmlEngine)

	return b
}

func (b *Bootstrapper) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.AppName,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

func (b *Bootstrapper) SetupErrorHandlers() {
	b.Application.OnAnyErrorCode(func(ctx iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  ctx.GetStatusCode(),
			"message": ctx.Values().GetString("message"),
		}

		if jsonOutput := ctx.URLParamExists("json"); jsonOutput {
			ctx.JSON(err)
			return
		}

		ctx.ViewData("Err", err)
		ctx.ViewData("Title", "Error")
		ctx.View("shared/error.html")
	})
}

// 启动计划任务服务
//func (b *Bootstrapper) setupCron() {
//	// 服务类应用
//	if conf.RunningCrontabService {
//		cron.ConfigueAppOneCron()
//	}
//	cron.ConfigueAppAllCron()
//}

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	b.SetupViews("./web/views")
	b.SetupSessions(24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetupErrorHandlers()
	b.Application.Favicon(StaticAssets + Favicon)

	//b.Application.setu

	return b
}

func (b *Bootstrapper) Listen(addr string, cfgs ...iris.Configurator) {
	b.Application.Run(iris.Addr(addr), cfgs...)
}