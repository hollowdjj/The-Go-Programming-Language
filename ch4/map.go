package ch4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func MapTest() {
	/*
		在Go语言中，map就是一个哈希表的引用。map类型可以写成map[K]V，其中键值K必须能够支持
		==比较运算符。内置的make函数可以创建一个map，也可以使用map的字面值语法创建并同时指
		定一些初始值。
	*/
	ages := make(map[string]int)
	ages1 := map[string]int{
		"alice":   34,
		"Charlie": 31,
	}
	fmt.Println(ages1)
	/*
		map的零值是nil，也就是没有引用任何哈希表。map上的大部分操作，例如查找、删除、len和range
		都可以完全地工作在nil值的map上。但是向一个nil值存入元素会导致panic异常，因为在向map存入
		元素之前，必须先创建map。在通过key值访问map时，如果key不存在，那么将得到value对应类型的
		零值。因此，为了区分一个已经存在的0(例如value的类型为0)，和一个不存在而返回的零值0，可以
		向下面这样测试。ok为一个布尔变量，true表示key存在，false表示不存在。
	*/
	if age, ok := ages["Bob"]; ok {
		fmt.Println(age)
	}
}

/*
	有时候我们需要一个map的key是slice类型，然而slice类型并不支持==运算符。我们可以通过定义一个
	辅助函数k来绕过这个限制。辅助函数k需要满足两个条件：1. k(slice)的返回值是一个可比较的类型；
	2. 只有当两个slice相等时，k(x)==k(y)成立。这种技术可以用于自定义某些类型的比较函数。
*/
var m map[string]int

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

//Dedup 通过map表示所有的输入行所对应的set集合
func Dedup() {
	input := bufio.NewScanner(os.Stdin)
	set := make(map[string]bool)
	for input.Scan() {
		line := input.Text()
		if !set[line] {
			set[line] = true
			fmt.Println(line)
		}
	}
	fmt.Println(set)
}

//CharCount 统计标准输入中每个Unicode码点出现的次数以及不同字节字符出现的次数
func CharCount() {
	counts := make(map[rune]int)
	var uftlen [utf8.UTFMax + 1]int //UTF-8编码，字符最常为4个字节，故用数组比map好
	invalid := 0                    //无效UTF-8字符的数量
	input := bufio.NewReader(os.Stdin)
	for {
		//ReadRune方法执行UTF-8解码并返回三个值：解码的rune字符的值，字符UTF-8编码后的长度，和一个错误值
		r, n, err := input.ReadRune()
		//可以预期的错误只有对应文件结尾的io.EOF
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		//如果输入的是无效的UTF-8编码的字符，r将是unicode.ReplacementChar表示的无效字符，并且编码长度为1
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		uftlen[n]++
	}
	fmt.Println("rune\tcount")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Println("\nlen\tocunt")
	for i, n := range uftlen {
		fmt.Printf("%d\t%d\n", i, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//Practice48 修改ChaCount，使用unicode.IsLetter等函数，统计字母、数字等Unicode中不同的字符类别
func Practice48() {
	counts := make(map[string]int)
	input := bufio.NewReader(os.Stdin)
	for {
		r, _, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Practice48: %v\n", err)
		}
		if unicode.IsLetter(r) {
			counts["letter"]++
		} else if unicode.IsDigit(r) {
			counts["digit"]++
		}
	}
	fmt.Printf("%v", counts)
}

//Practice49 报告输入文本中每个单词出现的频率
func Practice49() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	//在第一次调用Scan前调用，这样可以按照单词而不是按行取值
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		counts[word]++
	}
	fmt.Printf("%v", counts)
}
