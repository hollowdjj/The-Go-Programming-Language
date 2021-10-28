package ch1
//在Go中，一个目录中的一个或多个文件对应一个包(package)。也就是说，在同一个目录下的Go文件package后面的名字需相同
//import为用到了其他包，必须写在package之后，并且不能import没有使用的包，否则会导致编译错误
import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Echo1 打印命令行参数。Go中os包的Args切片提供了对程序命令行参数的访问。
// Args[0]为程序的名字，其余元素即为命令行参数
func Echo1() {
	var s,sep string      				//变量名在前，类型名在后
	for i := 1;i < len(os.Args);i++ {   //:=为短变量声明，Go中只有i++和i--，没有++i和--i
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func Echo2() {
	s,sep := " "," "				   //短变量声明会根据初始值来确定变量的类型，但只能用于声明局部变量，不能用来声明包一级的变量
	for _,arg := range os.Args[1:] {   //range创建一对值：索引和该索引的值。由于Go不允许未使用的局部变量，需要使用_空标识符占位
		s += sep + arg                 //多次拷贝，代价高昂
		sep = " "
	}
	fmt.Println(s)
}

func Echo3() {
	//使用join函数避免gc开销。
	fmt.Println(strings.Join(os.Args[1:]," "))
}

func Practice11() {
	fmt.Println(strings.Join(os.Args[0:]," "))
}

func Practice12() {
	for index,arg := range os.Args[0:]{
		fmt.Println(index," ",arg)
	}
}

func Practice13() {
	start1 := time.Now()
	Echo1()
	fmt.Printf("old version takes %dns\n",time.Since(start1).Nanoseconds())
	start2 := time.Now()
	Echo3()
	fmt.Printf("new version takes %dns\n" ,time.Since(start2).Nanoseconds())
}