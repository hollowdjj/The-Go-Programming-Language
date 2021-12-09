package ch10_Package

import (
	"fmt"
	"runtime"
)

/*
7. 工具
1) 工作区结构：go项目的目录结构如下所示。GOPATH对应的工作区目录有三个子目录，其中src用来存放源代码(.go文件)，bin用来存储
编译后的可执行程序，pkg用来保存编译后的包的目标文件。在src目录下，又有多个子目录，其中每一个子目录下都只能存放同一个包的源
文件。
   GOPATH/
      src/
       ch1/
         test.go
         ...
       ch2/
         9.1_Race.go
         ...
       ....
      pkg/
        linux_amd64/
        ...
      bin/
        ...
2) 下载包
用go get命名可以从互联网上下载go程序，go get -u可保证每个包都是最新版本

3)构建包
go build命名编译命令行参数指定的每个包。如果包是一个库(非main包)，则忽略输出结果；如果是main包，则会在当前目录下生成
一个可执行程序。
go install命令与go build命令的不同之处在于，其会保存包编译的结果而非直接丢弃。编译后的包的目标文件将会保存在GOPATH下
的pkg目录下。因为编译对应不同的操作系统平台和cpu架构，go install会在pkg下创建相应的目录保存包的目标文件，例如linux_amd64。
编译时也可指定64位和32位环境：GOARCH=386 go build ..... 。如果文件中包含注释： +build linux darwin表示只在Linux或
Mac OS X系统下才编译这个文件，而+build ignore则表示不编译这个文件

4) 包文档
go doc time       可查看time包的文档注释
go doc time.Since 可查看Since函数的注释
godoc -http :8000 可查看标准库源代码

5) 内部包
net/http
net/http/internal/chunked
net/http/httputil
net/url
一个internal包只能被和internal目录有同一个父目录的包导入。例如，net/http/internal/chunked内部的包，只能被net/http以及
net/http/httputil包导入(同一个父目录http)

6) 查询包
go list ...  查询工作区中的所有包
*/

func Cross() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
