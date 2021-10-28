package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//FetchAll 并发地请求多个URL
func FetchAll() {
	start := time.Now()
	ch := make(chan string)      //创建一个传递string类型参数的channel

	for _,url := range os.Args[1:] {
		//使用go关键字创建一个协程，fetch这个函数就在该协程中运行
		//同时，这些goroutine都使用同一个channel，即http请求的发送与响应的接收是异步的，而main中的消息打印是同步的
		go fetch(url,ch)
	}
	for range os.Args[1:] {
		//由于使用的是无缓存channel，若有一个网站一直不响应，不仅FetchAll会阻塞，所有其他成功响应了的goroutine也会阻塞
		fmt.Println(<-ch)       //channel用于在协程之间进行参数传递。若没有goroutine向ch发送数据，那么接收端阻塞
	}

	fmt.Printf("%.2f elapsed\n",time.Since(start).Seconds())
}

//fetch 向url发送Get请求。若成功，则丢弃响应报文并将耗时写入名为ch的channel中
func fetch(url string, ch chan<- string) {
	start := time.Now()
	response,err := http.Get(url)
	if err != nil {
		ch<- fmt.Sprint(err)
		return
	}
	nbytes,err := io.Copy(ioutil.Discard,response.Body)  //ioutil.Discard可视为垃圾桶，把不需要的数据扔进去
	response.Body.Close()
	if err != nil {
		ch<- fmt.Sprint(err)
		return
	}

	secs := time.Since(start).Seconds()
	ch<- fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)	//往channel中发送值。若没有goroutine在接收则阻塞
}