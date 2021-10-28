package ch8

import (
	"fmt"
	"os"
	"time"
)

func Countdown1() {
	fmt.Println("Commencing countdown.")
	//创建一个tick channel。程序会每个时间间隔d向该channel发送消息
	tick := time.Tick(1 * time.Second)
	for countdown := 10;countdown > 0;countdown-- {
		fmt.Println(countdown)
		<-tick //由于是无缓存channel，这里会阻塞，所以可以保证一个较为精确的倒计时
	}
	fmt.Println("launch!!!")
}


func Countdown2() {
	//启动一个goroutine负责从标准输入读取一个字节，读取到一个字节后，就向abort发送消息
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte,1))
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
	channel中发送值，这时就发生了goroutine泄露。因此，tick函数最好在程序的整个生命周期中都需要时再使用，其余
	情况下应使用time.NewTicker
	*/
	//tick := time.Tick(1 * time.Second)
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for countdown:= 10; countdown>0;countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick.C:
			//Do nothing
		case <- abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	fmt.Println("launch!!!")
}