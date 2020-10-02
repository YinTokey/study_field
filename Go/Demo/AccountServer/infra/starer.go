package infra

import "github.com/tietang/props/kvs"

const (
	KeyProps = "_conf"
)

type StarterContext map[string]interface{}

// 键值对操作
func (s StarterContext) Props() kvs.ConfigSource {
	p := s[KeyProps]
	if p == nil {
		panic("配置还没被初始化")
	}
	return p.(kvs.ConfigSource)
}

type Starter interface {
	//1.系统启动，初始化一些基础资源
	Init(StarterContext)
	//2. 系统基础资源的安装
	Setup(StarterContext)
	//3. 启动基础资源
	Start(StarterContext)
	//启动器是否可阻塞
	StartBlocking() bool
	//4. 资源停止和销毁
	Stop(StarterContext)

	//PriorityGroup() PriorityGroup

	Priority() int
}

//var _ Starter = new(BaseStarter)

func (s StarterContext) SetProps(conf kvs.ConfigSource) {
	s[KeyProps] = conf
}

type BaseStarter struct {
}

// 空实现
func (b *BaseStarter) Init(ctx StarterContext)  {}
func (b *BaseStarter) Setup(ctx StarterContext) {}
func (b *BaseStarter) Start(ctx StarterContext) {}
func (b *BaseStarter) StartBlocking() bool      { return false }
func (b *BaseStarter) Stop(ctx StarterContext)  {}

// register
type starterRegister struct {
	//nonBlockingStarters []Starter
	//blockingStarters 	[]Starter
	starters []Starter
}

func (r *starterRegister) Register(s Starter) {
	r.starters = append(r.starters, s)
}

func (r *starterRegister) AllStarters() []Starter {
	return r.starters
}

var StarterRegister *starterRegister = new(starterRegister)

func Register(s Starter) {
	StarterRegister.Register(s)
}

func SystemRun() {

	ctx := StarterContext{}

	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(ctx)
	}

	//2. 安装
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(ctx)
	}
	//2. 启动
	for _, starter := range StarterRegister.AllStarters() {
		starter.Start(ctx)
	}
}
