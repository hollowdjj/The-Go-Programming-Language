package ch8_Goroutines_Channels

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"time"
)

func RunClockServer() {
	port := flag.Int("p", 8080, "port number")
	flag.Parse()
	//监听localhost，默认端口为8080
	addr := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		//阻塞，直到一个新的连接被创建，此时返回一个net.Conn对象
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn) //每接受一个连接就开启一个goroutine进行处理
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	sys := runtime.GOOS
	var layout string
	switch sys {
	case "windows":
		layout = "15:04:05\r\n"
	case "linux":
		layout = "15:04:05\r\n"
	default:
		layout = "15:04:05\r\n"
	}
	for {
		//这里必须写成15:04:05，记忆方法1月2日下午3点4分5秒零六年UTC-0700，即1234567
		_, err := io.WriteString(conn, time.Now().Format(layout))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
