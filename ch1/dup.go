package ch1

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
Go语言格式化输出的动词：
%v    		按值本身的格式输出，万能动词，不知道用什么动词就用它
%+v	   		和%v一样，但在输出结构体时会添加字段名
%#v    		输出值的Go语法表示(先输出结构体名字，再输出字段类型+字段的值)
%t     		布尔值，true或者false
%c     		字符rune
%d     		十进制整数
%x,%o,%b 	十六进制，八进制，二进制整数
%O          八进制整数并添加ox前缀
%s          字符串
%q          带双引号的字符串
%T          变量的类型
%f    		浮点数，%.6f表示保留6位小数，如-123.456123
%e			科学计数法，%.6e表示有6位小数部分的科学计数法，如3.131314e+05
%g			根据实际情况采用%e或%f格式

bufio.NewScanner
bufio.NewScanner.Scan
os.Open
ioutil.ReadFile
*/

//Dup1 读取标准输入并打印重复出现的行
func Dup1() {
	counts := make(map[string]int)   		//创建一个key-type为string，value-type为int的哈希表(返回的是指针)
	//bufio包可以很方面的处理程序的输入和输出
	input := bufio.NewScanner(os.Stdin)		//Scanner类型用于读取输入并拆分成行或单词，类似于std::cin(需要使用ctrl+d来实现终止输入)
	for input.Scan() {						//Scan函数读入下一行并移除换行符。成功时返回true，反之返回false
		counts[input.Text()]++
	}

	for line,n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)  //类似C的printf
		}
	}
}

//Dup2 读取标准输入或者文件并打印重复出现的行
func Dup2() {
	counts := make(map[string]int)
	fileName := os.Args[1:]
	if len(fileName) == 0 {
		countLines(os.Stdin,counts)     //如果没有输入任何的命令行参数，那么就读取标准输入以获取文件名
	} else {
		for _,arg := range fileName {
			f,err := os.Open(arg)       //Open函数的第一个返回值时被打开的文件(*os.File)，第二个值是内置error类型的值
			if err != nil {
				fmt.Fprintf(os.Stderr,"dup2: %v\n",err)  //向标准错误流打印错误信息
				continue
			}
			countLines(f,counts)
			f.Close()
		}
	}

	for line,n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n",n,line)
		}
	}

}

// countLines 计算文件f的行数
func countLines(f* os.File,counts map[string]int) {       //go中只有值传递，然而make声明的map返回的是一个指针所以这里没问题
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

//Dup3 一次性读入文件的所有数据并根据需要进行拆分，而非一行行的读
func Dup3() {
	counts := make(map[string]int)

	for _,fileName := range os.Args[1:]{
		data,err := ioutil.ReadFile(fileName)  //ReadFile函数读取指定文件的全部内容
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _,line := range strings.Split(string(data),"\n") {  //根据换行符拆分字符串
			counts[line]++
		}
	}

	for line,n := range counts{
		if n > 1{
			fmt.Printf("%d\t%s\n",n,line)
		}
	}
}


func Practice14(){
	counts := make(map[string]string)   //key为有重复行的文件名 value为该文件的第一个重复行的内容
	fileName := os.Args[1:]
	if len(fileName) == 0 {
		processFile(os.Stdin,counts)
	} else {
		for _,arg := range fileName {
			f,err := os.Open(arg)      //Open函数的第一个返回值时被打开的文件(*os.File)，第二个值时内置error类型的值
			if err != nil {
				fmt.Fprintf(os.Stderr,"dup2: %v\n",err)  //向标准错误流打印错误信息
				continue
			}
			if processFile(f,counts) == true {
				fmt.Printf("%s has repeated lines with contents: %s\n",arg,counts[arg])
			}
			f.Close()
		}
	}
}

func processFile(f* os.File,counts map[string]string) bool {
	tempCounts := make(map[string]int)
	input := bufio.NewScanner(f)

	for input.Scan() {
		tempCounts[input.Text()]++
		if tempCounts[input.Text()] > 1 {
			counts[f.Name()] = input.Text()
			return true
		}
	}

	return false
}


