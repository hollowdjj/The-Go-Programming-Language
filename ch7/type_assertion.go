package ch7

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/*
类型断言是在一个接口值上进行的操作，语法为x.(T)，其中x必须是一个接口类型，T可以是一个具体类型或者接口类型。
T为具体类型时：x.(T)检查x的动态类型是否与T相同。如果检查成功，类型断言的结果是x的动态值。如果失败，则panic
T为接口类型时：x.(T)检查x的动态类型是否实现了T。如果检查成功，x的动态值不会被提取，返回值是一个T类型的接口值。
值得注意的是，无论T是一个接口类型还是一个具体的类型，只要x为一个nil值的接口，类型断言都会触发panic

类型断言返回两个结果可在断言失败时不触发panic，这时第二个结果是一个标识断言成功与否的布尔值
*/

func test() {
	var w io.Writer
	//注意os.Stdout是os包中一个类型为*os.File的包一级变量，而*os.File类型实现了 Write(p []byte) (n int, err error)方法
	w = os.Stdout

	//类型断言。w此时的动态类型为*os.File，断言成功，返回os.Stdout给f
	f := w.(*os.File)
	fmt.Println(f)

	//类型断言失败，触发panic，因为w的动态类型不是*bytes.Buffer
	c := w.(*bytes.Buffer)
	fmt.Println(c)

	//w的动态类型为*os.File，有Write，Read方法，实现了io.ReadWriter接口，断言成功。此时的返回值需要特别注意，其
	//返回的是一个动态类型和动态值不变的io.ReadWriter接口值(以前是一个io.Writer类型的接口值)。
	rw := w.(io.ReadWriter)
	fmt.Println(rw)

	//若类型断言成功，则ok为true否则为false并且此时不会触发panic
	f, ok := w.(*os.File)
	fmt.Println(ok)
}

//通过类型断言查询接口以实现不同的行为

/*
io.Writer的Write方法接收一个byte切片，而我们希望写入的是一个字符串，因而需要进行一个类型转换。
然而这样的转换会产生一个副本，造成资源浪费。
*/
func writerHeader(w io.Writer, contentType string) error {
	//if _,err := w.Write([]byte("Content-Type: ")); err != nil {
	//	return err
	//}
	//if _,err := w.Write([]byte(contentType)); err != nil {
	//	return err
	//}
	//
	//return nil
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}

	if _, err := writeString(w, contentType); err != nil {
		return err
	}

	return nil
}

/*
io.Writer的某些动态类型拥有一个字符串高效写入的WriteString方法。因此，可以使用类型断言查询
接口以优化上面的代码。
*/
func writeString(w io.Writer, s string) (n int, err error) {
	//用于类型断言
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}

	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) //避免了拷贝
	}

	return w.Write([]byte(s))
}
