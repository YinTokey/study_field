package main

import (
	"bufio"
	"fmt"
	"github.com/sony/sonyflake"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)


type User struct {
	ID             int64
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

var (
	globalID int
	idLocker sync.Mutex
	sf *sonyflake.Sonyflake

)

func main()  {

	FlakeInit()
	startServer()

}

func FlakeInit() {
	st := sonyflake.Settings{}

	flake := sonyflake.NewSonyflake(st)

	sf = flake
}

func NewGuid() int64 {
	//fmt.Println("start new guid ", sf)
	id, err := sf.NextID()
	result := int64(id)
	if err != nil {
		return 0
	}
	return result
}

func startServer() {

	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		log.Fatalf("listen error: %v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept error: %v\n", err)
			continue
		}
		// 开始goroutine监听连接
		fmt.Println("开始goroutine监听连接")
		go handleConn(conn)
	}

}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	user := &User{
		ID:             NewGuid(),
		Addr:           conn.RemoteAddr().String(),
		EnterAt:        time.Now(),
		MessageChannel: make(chan string, 8),
	}

	//启动写 conn 的协程
	go sendMessage(conn, user.MessageChannel)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Println("id:" + strconv.FormatInt(user.ID,10) + "  message:" + input.Text())
		//通过 chan 可以传递 message，客户端发来什么消息就回什么消息
		user.MessageChannel <- input.Text()
	}

	if err := input.Err(); err != nil {
		log.Println("读取错误：", err)
	}
}