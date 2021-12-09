package ch12_Reflect

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

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
反射由reflet包提供，它有两个重要的类型，Type以及Value，在test函数中有详细介绍。除此之外，还需要特别注意
Kind和Type的区别。Go中的类型(Type)指的是原生数据类型，如int string bool float32等以及使用type关键字定
义的用户类型，而Kind指的是对象归属的品种。Kind所表示的范围更大，类似于家用电器(Kind)和电视机(Type)间的对
应关系。举例说明：
*/

func test() {
	//函数reflet.TypeOf接受任意类型的interface{}类型并以reflect.Type形式返回其动态类型：
	t := reflect.TypeOf(3)  // 字面值3会隐式转换成一个接口类型，其动态类型为int，动态值为3
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"

	//reflect.TypeOf返回的是一个动态类型的接口值，是一个具体类型：
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // 打印"*os.File"，因为os.Stdout是一个*os.File类型的包一级变量

	//函数reflet.VauleOf接收一个任意类型的接口值，然后返回装载着该接口值的动态值的reflect.Value
	v := reflect.ValueOf(3) // v是一个reflect.Value类型的变量，其内保存有3这个值
	fmt.Println(v)          // "3"
	fmt.Printf("%v\n", v)   // "3"
	fmt.Println(v.String()) // "<int,Value>"
	//而对v调用Type方法，会返回具体类型对应的reflect.Type:
	t = v.Type()          // t是一个reflect.Type类型的变量
	fmt.Print(t.String()) // "int"

	//reflect.Value.Interface函数返回一个interface类型，装载着与reflect.Value相同的动态值
	x := v.Interface() //x是一个动态类型为int，动态值为3的接口值
	i := x.(int)       //类型断言，i即为x的动态值
	fmt.Printf("%d\n", i)
}
