
Goroutine 最直接的表面理解，就是 开一个 iOS gcd 的 global queue 线程来跑代码。（内部原理不一样，但是执行流程可以这么理解）



无缓冲通道 msg := make(chan  string)

有缓冲通道 msg :=  make(chan string, 2)   主要差异就在于构建时的第二个参数

```
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main () {

	done := make(chan bool,1)
	go worker(done)

	<- done
}
```

`<-done` 很关键，表示阻塞接收。它会打印 wroking.... done。 如果没有阻塞接收，那么打印 working.... , 程序就执行结束了

### 单向通道

通道可以定义成只读，或者只写

```
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
```

其中 pings 为只读通道，pongs 为只写通道





