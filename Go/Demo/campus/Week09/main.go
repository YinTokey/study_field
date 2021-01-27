package main

import (
	"bufio"
	"fmt"
	"github.com/sony/sonyflake"
	"log"
	"net"
	"strconv"
	"time"
)

type User struct {
	ID             int64
	Addr           string
	EnterAt        time.Time
	MessageChannel chan string
}

var (
	sf *sonyflake.Sonyflake
)

func main()  {
	// 初始化Id生成器
	FlakeInit()

	// 服务监听
	startServer()

}

func FlakeInit() {
	st := sonyflake.Settings{}

	flake := sonyflake.NewSonyflake(st)

	sf = flake
}

func NewGuid() int64 {
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
		fmt.Println("开始goroutine监听连接")
		go handleConn(conn)
	}

}

func sendMessage(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		str := fmt.Sprintf("reply from server : %s",msg)
		
		fmt.Fprintln(conn, str)

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

	go sendMessage(conn, user.MessageChannel)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		fmt.Println("from client id:" + strconv.FormatInt(user.ID,10) + "  message:" + input.Text())
		user.MessageChannel <- input.Text()
	}

	if err := input.Err(); err != nil {
		log.Println("读取错误：", err)
	}
}