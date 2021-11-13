package ch4

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func Array() {
	//在Go中，数组除了直接提供顺序初始化值序列，但是也可以指定一个索引和对应值列表的方式初始化
	nums := [...]int{0: 1, 1: 2, 2: 3}
	fmt.Println(nums) // "1 2 3 "

	//这里Currency的底层数据类型是int，因此也可以作为数组的索引
	type Currency int
	const (
		USD Currency = iota //美元
		EUR                 //欧元
		GBP                 //英镑
		RMB                 //人民币
	)
	symbol := [...]string{USD: "$S", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	//索引为5的元素被声明为-1，故数组长度为6
	r := [...]int{5: -1}
	//第一个元素为2，索引为3的元素为5，最后一个元素为2，故数组长度为5
	m := [...]int{2, 3: 5, 2}
	fmt.Println(r, m)
}

func Practice41() int {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	count := 0
	for i := 0; i < 32; i++ {
		if c1[i] != c2[i] {
			count++
		}
	}
	return count
}

func Practice42() {
	sha := flag.Int("s", 256, "SHA code")
	flag.Parse()
	text := flag.Args()
	for _, str := range text {
		switch *sha {
		case 256:
			fmt.Printf("%x", sha256.Sum256([]byte(str)))
		case 384:
			fmt.Printf("%x", sha512.Sum384([]byte(str)))
		case 512:
			fmt.Printf("%x", sha512.Sum512([]byte(str)))
		}
	}
}
