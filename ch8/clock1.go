package ch8

import (
	"io"
	"log"
	"net"
	"time"
)

func RunClockServer() {
	//监听127.0.0.1的8000端口
	listener, err := net.Listen("tcp","localhost:6688")
	if err != nil {
		log.Fatal(err)
	}

	for {
		//接收一个连接，阻塞
		conn,err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)  //每接受一个连接就开启一个goroutine进行处理
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		//这里必须写成15:04:05，记忆方法1月2日下午3点4分5秒零六年UTC-0700，即1234567
		_,err := io.WriteString(conn,time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}


