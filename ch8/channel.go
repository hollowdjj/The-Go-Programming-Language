package ch8

import (
	"fmt"
)

/*
channels是goroutine间的通信机制，可用于从一个goroutine向另一个goroutine发送值信息。每个channel都有一个特殊的类型
，也就是channels可发送数据的类型。一个可以发送int类型数据的channel一般写为chan int。使用内置的make函数，可以创建一
个channel。需要注意的是，channel和map一样，make函数返回的是一个底层数据结构的引用。

channel有不带缓存的和带缓存的两种类型。
ch = make(chan int)     //不带缓存的
ch = make(chan int,0)   //不带缓存的
ch = make(chan int,3)   //带缓存的

1. 不带缓存的channel
在一个不带缓存的channel上执行信息发送操作，将导致该goroutine阻塞，直到另一个goroutine在相同的channel上执行的信息接收
操作。当发送的值通过channel成功传输后，两个goroutine可以继续执行后面的语句。反之，若接收操作先发生，那么接收者goroutine
也会阻塞，直到发送者goroutine通过同一个channel执行了发送操作。不带缓存的channel的这种阻塞性质，可以用来实现两个goroutine
的同步。例如：当主goroutine必须要等待某个后台goroutine完成后才退出时，可以让后台goroutine在退出时向一个channel发送消息，
主goroutine则从该channel接受消息。

2. 带缓存的channel
内部持有一个元素队列，该队列的最大容量在创建channel时由传入的第二个参数确定。向一个带缓存的channel中发送信息就是向内部
缓存队列的尾部插入元素，接收操作则是从队列的头部删除元素。如果缓存区已满，那么发送操作将一直阻塞直到另一个goroutine执行
了接收操作。对应的，如果缓存区为空，那么接收操作将一直阻塞直到有另一个goroutine执行了发送操作。

内置的cap函数可以获取channel内部缓存区的容量，而len函数则获取的是channel内部缓存队列中有效元素的个数。
*/

func test() {
	//使用内置make函数创建一个可以发送int类型数据的channel，并返回其引用
	ch1 := make(chan int)
	ch2 := make(chan int)
	//channel对象可以使用==运算符进行比较
	if ch1 == ch2 {
		fmt.Println(ch1)
	}
	//发送值 接收值
	x := 10
	ch1 <- x  //发送
	x = <-ch1 //接收
	<-ch1     //接收但丢弃值

	/*
		向一个被关闭的channel发送消息将导致panic异常，而从一个被关闭的channel接收消息则不会导致panic异常，
		但此时接收到的值将是channel数据类型的零值。试图重复关闭channel，关闭一个nil值的channel也将导致panic
		异常。只有当需要告诉接收者goroutine所有数据都已经全部发送完时才需要关闭channel，因为go有gc
	*/
	close(ch1)
}

/*
串联的channel(pipeline)，n个channels可以用n-1个goroutine连接在一起，从而形成所谓的管道(pipeline)，如test1。
理想情况下，一个大的函数一般需要拆分成多个函数，这几个函数通过channel连接即可(单方向的channel)。
*/
func test1() {
	usual := make(chan int)
	square := make(chan int)

	//send
	go func() {
		for i := 0; i < 100; i++ {
			usual <- i
		}

	}()

	//receive and square
	go func() {
		for x := range usual {
			square <- x * x
		}
		close(square)
	}()

	//print
	for x := range square {
		fmt.Println(x)
	}
}

//chan<- int表示一个只发送int的channel
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
}

//<-chan int表示一个只接收int的channel
func squarer(out chan<- int, in <-chan int) {
	//循环从channel中读取值，当没有值可以读取时，自动退出循环
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}
