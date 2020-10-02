package base

import (
	"AccountServer/infra"
	"fmt"
	"github.com/tietang/props/kvs"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarer
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	fmt.Println("初始化配置")
}
