package ch4

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

/*
一个slice由三个部分组成：指针、长度和容量。其中指针指向第一个slice元素对应的底层数组元素地址。
值得注意的是，slice的第一个元素并不一定就是数组的第一个元素(因为slice可能引用的是一个数组的
的某一部分)。长度对应slice中的元素数量，容量一般是从slice的开始位置到底层数据的结尾位置。
*/

func SliceTest() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sept", 10: "Oct",
		11: "Nov", 12: "Dec"}

	//生成了一个新的slice Q2。Q2引用了months的4到7
	Q2 := months[4:7]
	fmt.Println(len(Q2), cap(Q2)) //"3 9"
}

/*
和数组不同的是，slice之间不能比较，因此不能用==操作符来判断两个slice是否含有全部相等的元素
不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相等([]byte)，对于其他
类型的slice，我们必须自己展开每个元素进行比较。slice唯一合法的比较操作是和nil比较。一个nil
值的slice并没有底层数组，其长度和容量都是0。值得注意的是，长度和容量都为0的slice却不一定是
一个nil值的slice，例如[]int{}或make([]int,3)[3:]。除此之外，所有函数都应该以相同的方式对
待nil值的slice和0长度的slice
*/
func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}

	/*
		Go中的for循环有三种形式：1. for init;condition;post { } 2. for condition { } 3. for { }
		range for也有三种形式： 1. for range slice { } 2. for k := range slice { }  3. for k,v := range slice{ }
	*/
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func equalStringSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func appendFunc() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		/*
			需要注意的是，通常我们是没法知道append函数的调用是否导致了内存的重新分配，即我们不能确定新的slice和旧的slice
			是否引用的是相同的底层数组。同样，我们也不能确定在原先的slice上的操作是否会影响到新的slice。因此，通常我们会
			将append返回的结果直接赋值给输入的slice变量。
		*/
		runes = append(runes, r)
	}
	//内置的append函数可以一次追加多个元素，甚至追加一个slice
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) //追加一个slice，必须有...
	fmt.Println(x)      //"1 2 3 4 5 6 1 2 3 4 5 6"
}

//nonempty 在原有的slice空间之上返回不包含空字符串的列表
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

//remove 删除slice中序号为i的元素
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:]) //用copy函数将后面的子slice向前依次移动一位
	return slice[:len(slice)-1]
}

//Practice43 重写reverse函数，用数组指针代替slice
func Practice43(array *[4]int) {
	for i, j := 0, 3; i < j; i, j = i+1, j-1 { //注意这里的pose部分不能是i,j = i++,j--或i++,j--
		array[i], array[j] = array[j], array[i]
	}
}

//Practice44 一次循环完成slice左旋n个数
func Practice44(slice []int, n int) {
	for i := 0; i < n; i++ {
		temp := slice[0]
		copy(slice[0:], slice[1:])
		slice[len(slice)-1] = temp
	}
}

//Practice45 原地完成消除[]string中相邻重复字符串的操作
func Practice45(strings []string) []string {
	prev, j := strings[0], 0
	for i := 1; i < len(strings); i++ {
		if strings[i] != prev {
			j++ //j是slice[0:i]中重复子序列的第二个元素
			strings[j] = strings[i]
		}
		prev = strings[i]
	}
	return strings[:j+1] //只取前j个元素
}

//Practice46 原地将一个UTF-8编码的[]byte类型的slice中相邻的空格替换成一个空格返回
func Practice46(b []byte) []byte {
	for i := 0; i < len(b)-1; i++ {
		if unicode.IsSpace(rune(b[i])) && unicode.IsSpace(rune(b[i+1])) {
			copy(b[i:], b[i+1:])
			b = b[:len(b)-1]
			i--
		}
	}
	return b
}

//Practice47 在不分配额外内存的情况下原地反转UTF8编码的[]byte
func Practice47(b []byte) []byte {
	l, r := 0, len(b)
	for l < r-1 {
		head, size1 := utf8.DecodeRune(b[l:])        //解码头部字符
		tail, size2 := utf8.DecodeLastRune(b[:r])    //解码尾部字符
		copy(b[l+size2:r-size1], b[l+size1:r-size2]) //将中间部分未处理的子串头部字符后面的字节向后移动size2-size1个字节
		copy(b[l:], []byte(string(tail)))            //填充新的头部
		copy(b[r-size1:], []byte(string(head)))      //填充新的尾部
		l += size2
		r -= size1
	}
	return b
}
