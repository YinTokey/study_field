package infra

import (
	"github.com/tietang/props/kvs"
	//"google.golang.org/appengine/log"
)

// 这个类似iOS的 appDelegate

type BootApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
}

// 构造方法
func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{conf: conf, starterContext: StarterContext{}}
	b.starterContext[KeyProps] = conf
	return b
}

func (b *BootApplication) Start() {
	// 1.  初始化starter
	b.init()
	// 2.  装载starter
	b.setup()
	// 3. 启动starter
	b.start()

}

func (b *BootApplication) init() {
	//log.Infof("Initializing starters")
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}

func (b *BootApplication) start() {
	for i, starter := range StarterRegister.AllStarters() {

		if starter.StartBlocking() {
			//如果是最后一个可阻塞的，直接启动并阻塞
			if i+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.starterContext)
			} else {
				//如果不是，使用goroutine来异步启动，
				// 防止阻塞后面starter
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}
	}
}
