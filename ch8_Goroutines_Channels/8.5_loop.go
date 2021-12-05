package ch8_Goroutines_Channels

import (
	"strings"
	"sync"
	"time"
)

/*
并发的循环
这一小节主要将的是，在循环体中使用并发所需要注意的一些技巧以及事项
*/

//ToUpper 将字符串中的所有小写字母转换成大写
func ToUpper(str string) string {
	time.Sleep(500 * time.Millisecond)
	return strings.ToUpper(str)
}

//Loop 在这个函数中，每一个字符串的处理是完全独立的，因此很容易想到采用并发来加速
//程序。然而，在下面这个版本的实现是错误的。原因在于，go语句会很快返回，因此，主
//goroutine很有可能在子goroutines之前退出。
func Loop(strs []string) {
	for _, str := range strs {
		go ToUpper(str)
	}
}

//Loop1 在这个版本的实现中，使用了一个channel来判断所有子goroutine是否都以退出。
func Loop1(strs []string) {
	ch := make(chan struct{})
	for _, str := range strs {
		go func(f string) {
			ToUpper(f)
			ch <- struct{}{}
		}(str) //闭包中引用的是变量的地址，所以这里必须显示传递参数
	}
	for range strs {
		<-ch
	}
}

//Loop2 当我们想要返回第一个转换完成的字符串时，可能会这样写。下面这段代码错误的地方在于
//函数返回后，后续没有清空channel的操作，因而会导致后续的goroutine一直阻塞，这种情况叫做
//goroutine泄露
func Loop2(strs []string) string {
	ch := make(chan string)
	for _, v := range strs {
		go func(f string) {
			ch <- ToUpper(f)
		}(v)
	}

	for range strs {
		return <-ch
	}

	return ""
}

//Loop3 解决Loop2函数中的问题的一个最简单的方法就是使用一个带缓存的channel
func Loop3(strs []string) string {
	ch := make(chan string, len(strs))
	for _, str := range strs {
		go func(f string) {
			ch <- ToUpper(f)
		}(str)
	}

	for range strs {
		return <-ch
	}

	return ""
}

//Loop4 Loop3中的解决方案只适用于我们实现知道循环次数的情况。当循环次数不确定时，我们需要
//在开启一个goroutine时使得一个计数器加1，而每个goroutine完成时都使计数器减1。这样，当计数
//器中的值为0时，才代表所有goroutine都以完成。满足这个需求的计数器由sync.WaitGroup提供。
func Loop4(str <-chan string) {
	var wg sync.WaitGroup
	for v := range str {
		//每开启一个goroutine就让wait group加1
		wg.Add(1)
		go func(f string) {
			//wait group减1
			defer wg.Done()
			ToUpper(f)
		}(v)
	}

	go func() {
		//阻塞直到wait group中的值为0
		wg.Wait()
	}()
}
