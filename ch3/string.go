package ch3

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func UnicodeTest() {
	/*
		在Go中，string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf8编码下占
		3个字节，而Go默认采用的是utf8编码。也就是说，在Go中，string的底层是一个代表了字符的utf8
		编码的byte数组。
	*/
	s := "Hello, 世界"
	//7 + 3 * 2 = 13(中文字符在UTF8中占3个字节) 对string而言，len函数返回的是“字节数”，而非“字符数”
	fmt.Println("len(s): ", len(s))
	//通过调用utf8.RuneCountInString(str)函数
	fmt.Println("utf8.RuneCountInString(s): ", utf8.RuneCountInString(s))
	//通过rune类型，对rune类型调用len，返回的就是真正的字符数量
	fmt.Println("rune: ", len([]rune(s)))
}

//Basename 硬编码实现删除文件的系统路径前缀以及后缀。例如Basename("a/b/c.go") = "c"
func Basename(s string) string {
	//去除前缀
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	//去除后缀
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//Basename1 使用strings库函数实现删除文件的系统路径前缀以及后缀
func Basename1(s string) string {
	s = s[strings.LastIndex(s, "/")+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//CommaInt 将一个表示整值的字符串，从后往前每隔三个字符插入一个逗号分隔符。例如“12345”处理后成为"12.345"
func CommaInt(s string) string {
	n := len(s)
	if n < 3 {
		return s
	}

	return CommaInt(s[:n-3] + "." + s[n-3:])
}

//CommaIntUsingBuffer 非递归版本的Comma函数，并使用bytes.Buffer代替字符串链接操作
func CommaIntUsingBuffer(s string) string {
	var buf bytes.Buffer
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Fprintf(&buf, "%c", s[i])
		if count++; count == 3 {
			buf.WriteByte('.')
		}
	}
	return buf.String() //待测试
}

//ByteSlice 展示string与byte slice间的相互转换
func ByteSlice() {
	/*
		字符串是一个包含只读字节的数组，一旦创建就不可更改。然而，一个字节的slice中的元素是可以任意修改的。
		并且，字符串和字节slice之间可以相互转换。
	*/
	s := "test"
	b := []byte(s)
	s1 := string(b)
	fmt.Println(s, b, s1)                //test [116 101 115 116] test
	fmt.Printf("%s\t%s\t%s\t", s, b, s1) //test test test
}

//PrintInts bytes.Buffer的使用
func PrintInts(nums []int) string {
	var buf bytes.Buffer
	//向bytes.Buffer中添加ASCII字符时，最好用WriteByte。
	buf.WriteByte('[')
	for i, n := range nums {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", n)
		//buf.WriteString(strconv.Itoa(nums[i]))
	}
	buf.WriteByte(']')
	return buf.String()
}
