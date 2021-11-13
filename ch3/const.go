package ch3

import (
	"fmt"
)

func TestConst() {
	//常量的批量声明
	const (
		a = 1
		b //如果没有指定值，那么会复制前面一个常量声明中右边的常量表达式
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"

	/*
		iota常量生成器，用于生成一组以相似规则初始化的常量。在第一个声明的常量所在的行，iota会
		被置为0，然后在每一个有常量声明的行iota的值会加一
	*/
	type Weekday int
	const (
		Sunday    Weekday = iota //0
		Monday                   //1
		Tuesday                  //2
		Wednesday                //3
		Thursday                 //4
		Friday                   //5
		Saturday                 //6
	)

	const (
		KB = 1000
		MB = 1000 * KB
		GB = 1000 * MB
		TB = 1000 * GB
		PB = 1000 * TB
		EB = 1000 * PB
		ZB = 1000 * EB
		YB = 1000 * ZB
	)
}
