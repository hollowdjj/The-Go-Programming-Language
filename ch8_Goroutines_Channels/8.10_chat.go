package ch8_Goroutines_Channels

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string //将client定义为一个只接收string类型数据的channel

var (
	entering = make(chan client) //客户连接
	leaving  = make(chan client) //客户断开连接
	messages = make(chan string) //消息
)

func broadcaster() {
	clients := make(map[client]bool) //所有已连接的用户
	for {
		select {
		case user := <-entering:
			clients[user] = true
		case user := <-leaving:
			delete(clients, user)
			close(user)
		case msg := <-messages:
			//广播消息
			for cli := range clients {
				go func(c client) {
					c <- msg
				}(cli)
				//select {
				//case cli <- msg:
				//	//Do nothing
				//default:
				//	continue
				//}
			}
		}
	}
}

func handleConnect(conn net.Conn) {
	user := make(chan string, 10) //这里的ch就是client
	go clientWriter(conn, user)

	who := conn.RemoteAddr().String()
	user <- "Me is " + who
	messages <- who + "has arrived"
	entering <- user

	//如果客户端静默时间超过了一定时间，则关闭连接
	duratin := 10 * time.Second
	tick := time.NewTimer(duratin)
	input := bufio.NewScanner(conn)
loop:
	for {
		select {
		case <-tick.C:
			conn.Close()
			break loop
		default:
			for input.Scan() {
				messages <- who + ": " + input.Text()
				tick.Reset(duratin)
			}
		}
	}

	leaving <- user
	messages <- who + "has left"
	conn.Close()
}

//向客户端写数据
func clientWriter(conn net.Conn, cli <-chan string) {
	//只有当cli被close后，循环才会终止
	for msg := range cli {
		fmt.Fprintln(conn, msg)
	}
}

//Chat 允许用户通过服务器向与该服务器连接的其他用户广播消息。
func Chat() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnect(conn)
	}
}
