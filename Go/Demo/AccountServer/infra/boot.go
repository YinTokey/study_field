package infra

import "github.com/tietang/props/kvs"

// 这个类似iOS的 appDelegate

type BootApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
}

// 构造方法
func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{conf:conf, starterContext: StarterContext{}}
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
	for _,starter := range star
}

func (b *BootApplication) setup() {

}

func (b *BootApplication) start() {

}


