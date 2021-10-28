package ch1

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

/*
知识点：
rep,err = http.Get(url)         发送get请求，并接收服务器的响应
ioutil.ReadAll(rep.Body)        读取响应报文的全部内容
rep.Body.Close()                关闭Body流，防止资源泄露
os.Exit(1)					 	退出进程
io.Copy(dst,src)  				从src中读取内容，并将读取的结果写入到dst中，可避免拷贝
strings.HasPrefix(s,pre)        判断字符串s是否含有前缀pre
*/


func Fetch() {
	for _,url := range os.Args[1:] {
		//发送get请求并接收http响应报文。若成功，可以从response这个结构体中获取服务器的响应
		response,err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch %v\n",err)
			os.Exit(1)				  //终止进程
		}
        //ReadAll函数读取响应报文的所有数据。最后需要关闭Body流，防止资源泄露
		b,err := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}
}

func Practice17() {
	for _,url := range os.Args[1:] {
		response,err := http.Get(url)         //发送get请求并接收http响应报文
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch %v\n",err)
			os.Exit(1)				  //终止进程
		}

		//io.Copy(dst,src)从src中读取内容，并将读取的结果写入到dst中，可避免拷贝
		num,err := io.Copy(os.Stdout,response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
			os.Exit(1)
		}

		fmt.Printf("\n\n Total %d bytes received\n",num)
	}
}

func Practice18() {
	for _,url := range os.Args[1:] {
		if !strings.HasPrefix(url,"http://") {
			url = "http://" + url
		}

		response,err := http.Get(url)         //发送get请求并接收http响应报文
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch %v\n",err)
			os.Exit(1)				  //终止进程
		}
		num,err := io.Copy(os.Stdout,response.Body)
		response.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
			os.Exit(1)
		}

		fmt.Printf("\n\n Total %d bytes received\n",num)
	}
}

func Practice19() {
	for _,url := range os.Args[1:] {
		response,err := http.Get(url)         //发送get请求并接收http响应报文
		if err != nil {
			fmt.Fprintf(os.Stderr,"fetch %v\n",err)
			os.Exit(1)				  //终止进程
		}

		fmt.Printf("%s\n",response.Status)
	}
}

