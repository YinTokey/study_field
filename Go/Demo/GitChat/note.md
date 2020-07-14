### 主要逻辑
服务端响应请求，模板引擎处理完毕后，直接下发HTML实现服务端渲染。

### 语法笔记
go语言引入一个新的赋值符号 “:=”，冒号后面紧跟一个等号，用于定义变量并直接初始化。比如下面式子等价。

```
tmpNewVar := 100
var tmpNewVar int = 100
var tmpNewVar = 100
```
“:=” 左边要是不存在的变量，而且不需要前面再加var了。



# Chapter 3
为什么请求会产生404错误？  服务端收到请求后，但是没有相应的处理器（handler）来处理这个请求，所以相应了404给客户端。

#### 处理器和处理器函数
二者存在一一对应的等价关系

从定义上说处理器函数就是和处理器拥有相同行为的函数。


以下是处理器
```
type HelloHandler struct {}

func (h *HelloHandler) ServeHTTP(writer http.ResponseWriter,request *http.Request){
	print("xx")
}

type WorldHandler struct {

}

func (h *WorldHandler) ServeHTTP(writer http.ResponseWriter,request *http.Request){
	print("aaaa")
}


func main() {

	heHandler := HelloHandler{}
	woHandler := WorldHandler{}

	server := http.Server{
		Addr:    "0.0.0.0:8080",
	}

	http.Handle("/hello",&heHandler)
	http.Handle("/world",&woHandler)

	server.ListenAndServe()
}

```

以下是对应的处理器函数

```
func hello(writer http.ResponseWriter,request *http.Request){
	print("xx")
}

func world(writer http.ResponseWriter,request *http.Request){
	print("aaaa")
}


func main() {


	server := http.Server{
		Addr:    "0.0.0.0:8080",
	}

	http.HandleFunc("/hello",hello)
	http.HandleFunc("/world",world)

	server.ListenAndServe()
}

```

从主观上来说，使用处理器函数更简洁易懂。二者之间存在转化关系，处理器函数 ，可以被转化为处理器（结构体）。
也就是说处理器函数是一种创建处理器的简便方法。但是如果是更复杂的场景，需要做模块化，那么还是使用处理器更强大一点。

`横切关注点`（cross-cutting concern）: 日志记录，安全检查，错误处理

Go 语言可以实现 函数作为另一个函数的参数或者返回值。


`ServeMux`：http的请求多路复用器，是一个结构体，里面有 url-handler 映射表
`DefaultServeMux`: serveMux 的一个实例，它是 net/http 标准库写好的一个默认实例。就像单例对象，名字都定好了。当然本质上不能当它是单例。

ServeMux本质上也是处理器。

ServeMux的缺点是不能做URL模式匹配，特别不适合处理带参数的GET请求。`HttpRouter`更适合处理这种。

`包管理`
通过执行命令，就可以把包从github下载到本地 go/src/github 默认包统一存储路径
```
go get github.com/julienschmidt/httprouter
```
当然不止github, 其他地方也可以
```
go get ”golang.org/x/net/http2” 
```

# Chapter 4
请求的表单，就是键值对形式存储数据在HTML中。标准库就支持解析表单了
http.ResponseWriter 用它来响应客户端

# Chapter 5
```
t ’ err := template.ParseF工les (”tmpl.html”) 
```
表示解析模板，结果赋值给 t, 如果解析发生错误，结果赋值给err。 判断err 是否有值，来判断解析成功还是失败。这是 Go 常见的一种写法

上面等价于
```
t : = template .New(”tmpl.html”) t , err : = t.ParseFiles (”tmpl.html”) 

```

有点类似于iOS的便利构造方法。

# Chapter 6
文件存储格式  CSV,  gob。 
gob是Go特有的，能将内存中的数据结构转换成二进制形式，实现数据持久化存储。有点类似iOS的归档解档。

`ORM` 对象关系映射，将对象转换为数据库存储。从iOS角度来说，就是对FMDB进一步封装，让model可以便利地存储到数据库中。

# Chapter 7
`SOAP`: XML
`RESTFUL` : JSON。 显然SOAP是比较笨重过时的

# Chapter 10
项目部署方法，更推荐 Docker 和独立服务器，而不是 Heroku, GAE








