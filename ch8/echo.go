package ch8

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn,shout string,delay time.Duration) {
	fmt.Fprintln(c,"\t",strings.ToUpper(shout))  //全大写
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",shout)					 //正常
	time.Sleep(delay)
	fmt.Fprintln(c,"/t",strings.ToLower(shout))  //全小写
}

func handleConn_(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c,input.Text(),1 * time.Second)
	}

	c.Close()
}
