package ch4

import (
	"fmt"
	"time"
)

type Employee struct {
	ID            int
	Name, Address string    //相邻的成员类型如果相同的话，可以被合并到一行
	DoB           time.Time //结构体成员的输入顺序是有意义的。不同的顺序将导致不同的结构体类型
	Position      string
	Salary        int
	ManagerID     int
	password      string //结构体成员名大写开头表示可导出的，小写开头表示不可导出的
}

/*
一个结构体中的成员可以是另一个结构体，这种情况被称为结构体嵌套。然而，结构体嵌套就会导致访问每个
成员变得繁琐，例如：
type Point struct {
	X,Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	shape Circle
	Spokes int
}

var w Wheel
w.Circle.Center.x = 8
*/

/*
Go语言的一个特性可以让我们只声明一个成员对应的数据类型而不指定成员的名字。这类成员被称为
匿名成员。匿名成员的数据类型必须是一个命名的类型或指向一个命名类型的指针。得益于匿名嵌入
的特性，我们可以直接访问叶子属性而不需要给出完整的路径。除此之外，外层的结构体不仅仅是获
得了匿名成员类型的所有成员，而且也获得了该类型导出的全部的方法。
*/

type Point struct {
	X, Y int
}

type Circle struct {
	Point  //匿名成员，说Point类型被嵌入到了Circle结构体中
	Radius int
}

type Wheel struct {
	Circle //匿名成员，说Circle类型被嵌入到了Wheel结构体中
	Spokes int
}

func TestStruct() {
	//匿名变量在访问叶子属性时带来的便利性
	var w Wheel
	w.X = 8
	w.Y = 9
	//在匿名变量存在时，需要特别注意结构体字面值的写法
	w = Wheel{Circle{Point{8, 9}, 10}, 20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{8, 8},
			Radius: 5,
		},
		Spokes: 20,
	}

	//%#v会额外打印结构体中每个成员的名字
	fmt.Printf("%#v", w)
}
