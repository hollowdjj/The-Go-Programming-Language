package ch12_Reflect

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

func testReflect() {
	/*
		reflect.Type(a)接受任意的interface{}类型，并以reflect.Type的形式返回其动态类型。
		也就是说，reflect.Type函数的返回值为一个带有接口值a动态类型的reflect.Type对象。
	*/
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) //w是一个接口，其动态类型为*os.File，故打印值为*os.File
	t := reflect.TypeOf(3)
	fmt.Println(t) //int
	/*
		reflect.Value(a)同样是接受任意的interface{}类型。其返回值是一个装有接口值a动态值的
		reflect.Value对象。对reflect.Value对象调用Type方法，将返回具体类型所对应的reflect.Type
	*/
	v := reflect.ValueOf(3)
	fmt.Println(v)          //3
	fmt.Println(v.String()) //<int Value>
	fmt.Println(v.Type())   //int
	/*
		reflect.Value.Interface是reflect.ValueOf的逆操作。返回的是一个接口类型，装载着与reflect.Value
		相同的具体值。
	*/
	d := reflect.ValueOf(3) //v是一个reflect.Value类型的对象
	x := d.Interface()      //x是一个interface{}，其动态类型为int，动态值为3
	i := x.(int)            //类型断言成功，返回动态值，即i的值为3
	fmt.Println(i)
}

/*
fmt.Fprintf函数可以对任意类型的值进行格式化并打印，甚至支持用户自定义的类型。
下面尝试用Sprint函数实现一个类似的功能。
*/

func Sprint(x interface{}) string {
	//传入的类型必须定义了String() string才支持打印，所以这里需要使用类型断言进行判断
	type stringer interface {
		String() string
	}

	//显然，在这里，我们不肯能列出所有类型，这也是为什么需要反射的原因
	switch x := x.(type) { //x.(type)也是一种类型断言，返回的是x的类型
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
	}

	return ""
}

/*
利用反射实现任意类型值得格式化打印。需要注意的是，在go中由于type关键字，因此我们可以创造无穷无尽的类型。
然而，所有这些类型的底层数据类型的数量确实有限的。
*/

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	//v.Kind返回的是v所装载的值的底层类型
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		//返回v的值，但类型为Int64。若v.Kind()不是以上任意一种类型的话，会panic
		//strconv.FormatInt(i int64,base int)是指将i转换成“base"进制的数。
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	//......
	default:
		return v.Type().String() + "value"
	}
}

/*
通过reflect.Value修改值，即可取地址的reflect.Value
*/
func reflectChangeValue() {
	x := 2
	/*
		所有通过reflect.ValueOf(x)返回的reflect.Value都是不可取地址的。但是对于变量d，他是c解引用
		方式生成的，指向另一个变量，因而是可以取地址的。我们可以通过调用reflect.ValueOf(&x).Elem()
		来获取任意变量x对应的可取地址的value。通过调用reflect.Value的CanAddr方法可以判断其是否可以
		被取地址。
	*/
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x) //只有参数为指针时，返回的reflect.Value才是可取地址的
	d := c.Elem()

	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())
	fmt.Println(d.CanAddr())

	/*
		要从一个变量对应的可取地址的reflect.Value来访问变量需要三个步骤：
		1. 调用Addr方法，它返回一个Value，里面保存了指向变量的指针
		2. 在Value上调用Interface()方法，它返回的是一个interface{}，里面包含有指向变量的指针(其实就是把reflect.Value转换成interface{})
		3. 在interface上使用类型断言，即可获取接口值的动态类型，也即指向变量x的指针
	*/
	y := 2
	v := reflect.ValueOf(&y).Elem()
	px := v.Addr().Interface().(*int)
	*px = 3
	fmt.Println(y)
	/*
		或者，不使用指针，而是直接在一个可取地址的reflect.Value的Set方法直接
		更新变量的值
	*/
	v.Set(reflect.ValueOf(4))

}
