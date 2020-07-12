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


