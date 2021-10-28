package ch1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int
var cycle  = 5

func Server() {
	http.HandleFunc("/",handler)  				//将所有发送到/路径下的请求交由hander1函数处理
	http.HandleFunc("/count",counter)
	http.HandleFunc("/print",printHttpMsg)
	http.HandleFunc("/update",changeCycle)
	//log.Fatal等同于Print后接一个Exit(1)
	//ListenAndServe监听welcoming socket并在有请求到来时调用handler执行响应的操作
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

//handler 访问次数加一并返回URL
func handler(w http.ResponseWriter,r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,"URL.PATH = %q\n",r.URL.Path)   //%q表示是带双引号的字符串和带单引号的字符
}

//counter 返回访问次数
func counter(w http.ResponseWriter,r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w,"count=%d",count)
	mu.Unlock()
}

// printHttpMsg 打印http请求报文的头部行和首部行
func printHttpMsg(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"%s %s %s\n",r.Method, r.URL,r.Proto)    //头部行
	for k,v := range r.Header {                                   //首部行
		fmt.Fprintf(w,"Header[%q] = [%q]\n",k,v)
	}
	fmt.Fprintf(w,"Host = %q\n",r.Host)
}

//changeCycle 更改cycle的值
func changeCycle(w http.ResponseWriter,r *http.Request) {
	val,err := strconv.Atoi(r.URL.Query().Get("cycles"))
	if err != nil {
		fmt.Fprintf(w,"invalid number for cycles\n")
	}
	cycle = val
	fmt.Fprintf(w,"The current value of cycle is %d!",cycle)
}

