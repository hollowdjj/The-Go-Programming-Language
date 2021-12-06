package ch8_Goroutines_Channels

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func Countdown1() {
	fmt.Println("Commencing countdown.")
	//创建一个tick channel。程序会每个时间间隔d向该channel发送消息
	tick := time.Tick(1 * time.Second)
	//这里稍微解释一下这个循环。可以预期的是，打印、countdown--以及判断countdown > 0
	//的时间是肯定远远小于1秒的。再加上是一个无缓存channel，这里会阻塞，所以可以保证一
	//个较为精确的倒计时。
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	fmt.Println("launch!!!")
}

func Countdown2() {
	//启动一个goroutine负责从标准输入读取一个字节，读取到一个字节后，就向abort发送消息
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	/*
		为了实现用户按下return键就停止，每一次计数循环的迭代都需要等待两个channel中的其中一个返回事件。也就是
		说，需要同时监听两个channel。这个需求就很像linux里面的多路复用。在go中，提供了select关键字实现多路复用。
	*/
	fmt.Println("Commencing countdown.")
	/*
		select会等待case中有能够执行的case时去执行。在执行一个case期间，其余case是不会执行的。如果有多个case同时
		就绪，select会随机地选择一个执行。需要注意的是，time.Tick创建的goroutine在函数结束过后依然存活，会继续向
		channel中发送值，这时就发生了goroutine泄露。因此，tick函数最好在程序的整个生命周期中都需要时才使用，其余
		情况下应使用time.NewTicker。除此之外，还可以使用time.After函数，该函数会立即返回一个channel，并开启一个
		goroutine，且在经过特定时间后向该channel发送一个值。
	*/
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		/*
			select语句中的每一个case代表一个通信操作，即在某一个channel上进行发送或接收操作。select会等待case中有能够
			执行的case时才去执行，其他时候select阻塞。因此，一个没有任何case的select语句写作select{}会一直阻塞下去。也
			就是说，select是会阻塞，但是一旦执行了某一个case后，select语句就执行完成了。因此，如果要实现多次轮询，还需要
			配合一个while循环。
		*/
		select {
		case <-tick.C: //time.After(10 * time.Second)
			//Do nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		default:
			//default语句可以用来实现当其他的操作都不能够马上被处理时，程序需要执行哪些逻辑
			//Do nothing
		}
	}

	fmt.Println("launch!!!")
}

func EchoAdvanced() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleEchoRequest(conn)
	}
}

func handleEchoRequest(c net.Conn) {
	input := bufio.NewScanner(c)
	tick := time.NewTicker(10 * time.Second)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			fmt.Fprintln(c, "Since you are silent for more than 10s, the connection is shut down!")
			return
		default:
			if input.Scan() {
				text := input.Text()
				fmt.Println(text)
				fmt.Fprintln(c, "Echo: ", text)
				tick.Reset(10 * time.Second)
			}
		}
	}
}
