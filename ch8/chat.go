package ch8

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string    //将client定义为一个只接收string类型数据的channel

var (
	entering = make(chan client)       //客户连接
	leaving = make(chan client)        //客户断开连接
	messages = make(chan string)       //消息
)

func broadcaster() {
	clients := make(map[client]bool)  //所有已连接的用户
	for {
		select {
		case user := <- entering:
			clients[user] = true
		case user := <-leaving:
			delete(clients,user)
			close(user)
		case msg := <-messages:
			for user := range clients {
				user <- msg
			}
		}
	}
}

func handleConnect(conn net.Conn) {
	ch := make(chan string)    //这里的ch就是client
	go clientWriter(conn,ch)

	who := conn.RemoteAddr().String()
	ch <- "Me is " + who
	messages <- who + "has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving  <- ch
	messages <- who + "has left"
	conn.Close()
}

//向客户端写数据
func clientWriter(conn net.Conn,ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn,msg)
	}
}
// 允许用户通过服务器向与该服务器连接的其他用户广播消息。
func Chat() {
	listener,err := net.Listen("tcp","localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnect(conn)
	}
}

