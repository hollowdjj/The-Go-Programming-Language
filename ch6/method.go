package ch6

import (
	"fmt"
	"image/color"
	"math"
	"time"
)

type Point struct {
	X, Y float64
}

//Distance 传统的函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Distance 定义了一个专属于Point类型的方法。其中p被称为接收器，推荐以类型首字母命名。
//但需要注意的是，我们不能声明一个名为X或Y的方法，因为其和结构体成员重名了。
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

//ScaleByPointer 基于指针对象的方法。一般的，我们会约定，如果Point这个类有一个指针作为接收器方法
//那么所有Point的方法都必须有一个指针接收器。
func (p *Point) ScaleByPointer(factor float64) {
	fmt.Println("call ScaleByPointer")
	p.X *= factor
	p.Y *= factor
}

func (p Point) ScaleByValue(factor float64) {
	fmt.Println("call ScaleByValue")
	p.X *= factor
	p.Y *= factor
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

/*
需要注意的是，如果类型本身是一个指针，那么我们不能对其定义任何方法，例如：
type P * int
func (p P) 	test() {...}    //error
func (p *P) test() {...}    //error
*/

type ColoredPoint struct {
	Point
	Color color.RGBA
}

func TestMethod() {
	/*
		关于接收器有以下两点需要注意：
		1. 编译器允许接收器和类型和调用方法的类型不一致，会自动将其转成接收器的类型(即取地址再调用和解引用再调用)
		2. 如果接收器是值类型的，那么在该方法中，操作的都是拷贝的那个对象，而非调用该方法的那个对象。指针类型的话
		   操作的就是调用它的那个对象。
	*/
	p := Point{
		X: 4,
		Y: 2,
	}
	p.ScaleByValue(2)
	fmt.Println(p) //{4,2} 接收器为值类型，那么操作的是p对象的一个拷贝
	p.ScaleByPointer(2)
	fmt.Println(p) //{8,4} 接收器为指针类型，那么操作的就是对象本身
}

//TestMethod1 通过嵌入结构体来扩展类型
func TestMethod1() {
	/*
		结构体内嵌可以让我们直接访问内嵌结构体的成员而不需要显示的指明它。对于内嵌类型的方法，我们也可以有类似的操作。
		即，我们可以把ColoredPoint类型当做接收器来调用Point里的方法，即使ColoredPoint中并没有声明这些方法。当编译器
		解析一个方法的调用时，它会首先找直接定义在这个类型里的方法，然后再递归的去找匿名类型的同名方法。
	*/
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	p := ColoredPoint{Point{1, 1}, red}
	q := ColoredPoint{Point{5, 4}, blue}
	//注意，我们虽然可以用ColoredPoint作为Point类型Distance方法的接收器，但方法的参数类型仍然是Point，因此需要
	//显示选择
	fmt.Println(p.Distance(q.Point)) // 5
	p.ScaleByPointer(2)
	q.ScaleByPointer(2)
	fmt.Println(p.Distance(q.Point)) //10
}

type Rocket struct{}

func (r *Rocket) Launch() {
	fmt.Println("Rocket launch!!!")
}

//TestMethod2 方法值
func TestMethod2() {
	/*
		我们经常选择一个方法，并且在同一个表达式里执行，比如常见的p.Distance()形式。然而，实际上，这个函数调用可以分两
		步进行。p.Distance叫做“选择器”，返回一个“方法值”。这个方法值将方法绑定到了特定接收器变量上，可以直接调用。也就是
		说，“方法值是一个绑定了特定接收器的方法”。
	*/
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q)) // "5"
	/*
		方法值在一个包的API需要一个函数值，并且该API的调用者希望传入的函数绑定了一个特定的对象时会非常有用。因为这样可以避免
		写一个匿名函数。
	*/
	r := new(Rocket)
	time.AfterFunc(10*time.Second, func() { r.Launch() })
	time.AfterFunc(10*time.Second, r.Launch)
}

//TestMethod3 方法表达式
func TestMethod3() {
	p := Point{1, 2}
	q := Point{4, 6}
	/*
		当T是一个类型时，方法表达式可能会写作T.f或(*T).f，会返回一个函数值。这种函数值会将其第一个参数用作接收器
	*/
	distance := Point.Distance
	fmt.Println(distance(p, q)) //此时第一个参数为接收器
}
