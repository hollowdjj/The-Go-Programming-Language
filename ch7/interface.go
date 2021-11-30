package ch7

import (
	"fmt"
	"io"
	"os"
)

/*
接口类型其实就是一个方法集合，一个实现了这些方法的具体类型就是这个接口类型的实例。一个类型如果
实现了接口的所有方法，那么这个类型就实现了这个接口。
*/

func TestInterface() {
	var w io.Writer
	/*
		需要注意的是，*T和T是两种不同的类型。若只实现了*T类型的String方法，那么也只有*T类型
		实现了fmt.Stringer接口，T类型则没有。
	*/
	w = os.Stdout //*os.File类型实现了io.Writer接口

	//可以将任意一个值赋值给一个空接口类型
	var any interface{}
	any = true
	any = "21"
	fmt.Println(w, any)

	/*
		接口的值，即接口值由两部分组成，一个具体的类型和那个类型的值，称为接口的动态类型和动态值
		接口值是可以使用==和!=进行比较的。两个接口值相等仅当它们都是nil值或者它们的动态类型相同
		并且动态值也根据这个动态类型的==操作相等。
	*/
	var w1 io.Writer //w1是一个零值接口，动态类型和动态值均为nil
	w1 = os.Stdout   //w1的动态类型为*os.File，动态值为*os.Stdout
	w1.Write([]byte("hello"))
	w1 = nil

}
