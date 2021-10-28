package ch2

import (
	"flag"
	"fmt"
	"strings"
)

/*
这一个例子主要讲解的是Go标准库中的flag包。命令行程序中，经常需要输入各种各样的命令行标志参数，例如
-h(-help)。flag包就提供了解析用户输入的命令行参数的功能。例如，flag.Bool用于创建一个布尔类型的命令
行标志参数变量，函数的三个参数依次为“命令行标志参数的名字”，“默认值”以及“描述信息”，返回值是一个指针
，指向存有用户输入的该命令行标志参数值的变量地址。
*/
var n = flag.Bool("n",false,"omit trailing newline")
var sep = flag.String("s"," ","separator")


func Echo() {
	flag.Parse()         						//解析命名行参数。解析出的值就会被保存到相应的命令行标志参数变量中
	fmt.Print(strings.Join(flag.Args(),*sep))   //flag.Args用于访问非标志参数
	if !*n {
		fmt.Println()
	}
}

/*
TestNew
Go中new是一个内置函数而非关键字。并且与C++中的new完全不同的是，go中的这个内置new函数只是创建一个T类
型的变量，初始化为T类型的零值，然后返回变量的地址。至于变量在堆区还是栈区是完全不可控的。
*/
func TestNew() {
	p := new(int)
	fmt.Print(*p)
}

var global *int
func f() {
	x := 1
	/*
	在这里，变量x会编译器分配在堆区。因为，在函数执行完毕后，变量x还是能够通过包一级的指针变量global
	找到。用Go语言的术语说，这个局部变量x从函数f中“逃逸”了。所以，一个变量的有效生命周期只取决于是否
	可达，从而一个局部变量的生命周期可能超出局部作用域。因此，我们需要尽量避免将一个指向短生命周期对象
	的指针保存到具有长生命周期的对象中。
	*/
	global = &x
}

func g() {
	p := new(int)
	*p = 1     //*y没有从函数中逃逸，因此编译器可以选择将其分配在栈区，也可以选择分配在堆区。
}