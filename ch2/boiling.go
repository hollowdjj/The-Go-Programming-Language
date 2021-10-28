/*
Package ch2
在go中，每一个go文件都需以package $name$开头以表明这个go文件属于哪一个包。并且，同一目录下的所有go文件都必须
属于同一个package。在go文件中，可以通过import导入其他package，但必须写在packgae关键字之后。

go语言一共有4种主要类型声明语句: var const type func，分别对应变量 常量 类型 函数。这些类型也被称为包一级类型，
其声明顺序无关紧要，但在函数内部的名字就必须要先声明后引用。

1. 变量var
变量声明的一般语法为： var 变量名 类型 = 表达式。其中“类型”和“=变量名”可以省略其中的一个。同时，可以在一个声明语句
中声明相同类型或不同类型的变量，如 var b,f,s = true,1.2,"str"；也可以通过函数调用的返回值初始化一组变量，例如
var f,err = os.Open(name)。

var声明一般用于需要显示指定变量类型的地方，或者变量在稍后会被重新赋值而初始值无关紧要的地方。局部变量的声明和初始化
常常采用“短变量声明:=”。需要注意的是，:=是一个声明语句，而非赋值语句，因此，在短变量声明语句中必须“至少声明一个变量”
而另外的变量可以是赋值行为。同时，若:=声明的是一个和包一级变量同名的变量，此时发生的是声明行为而非赋值行为。

在go中，返回局部变量的地址是安全的。这是因为编译器会自动进行“逃逸分析”，当发现变量的作用域没有跑出函数范围时，将它分配
在栈上，否则分配在堆上。go中的new(T)函数用于创建一个T类型的匿名变量，初始化为T类型的零值，然后返回该变量的地址。注意，
new只是一个预定义的函数，并不是一个关键字。并且，变量是声明在栈上还是堆上是由编译器决定的，而非由var或者new声明变量的方
式决定。

包一级变量的生命周期和程序的运行周期相同，而局部变量的生命周期是从其声明开始，到不再使用为止，gc会进行自动回收。如果将
指向短生命周期对象(局部变量)的指针保证到具有长生命周期的对象中(包一级变量)，会使得局部变量从局部作用域中“逃逸”了，从而
影响gc，进而可能影响程序性能。

2. 类型type
类型声明语句用于使用现有的底层结构来创建一个新的类型名称，语法为 type 类型 底层类型，例如 type Celsius float64需要注
意的是，若多个type声明出来的类型采用的是同一个底层类型，他们之间也是不兼容的，不能相互比较或混在同一表达式中进行运算。
底层类型转换为type对应的声明需要使用Celsius(t)的形式，这种形式并非函数调用，而是类型转换，它并不会改变值，只是使改变了
值的类型。

同时，我们还可以为type声明的类型定义方法(有点类似成员函数？)。例如，fun (c Celsius) String() string {............}
上面这个函数的意思就是，为Celsius类型的变量定义了一个名为String的方法，c就指代调用这个方法的Celsius变量。

3. 常量const
常量的声明与变量相同，只是将var替换成const。在批量声明常量时，可进行一定的省略，如：
const(
	a = 1     //1
	b         //1
	c = 2     //2
	d         //2
)
以上特性导出了iota常量生成器，如
type Weekday int
const (
	Sunday Weekday = iota   //第一个声明的常量所在的行，iota会被置0，后续每一行iota依次加1
	Monday					//1
	Tuesday					//2
	Wednesday				//3
	Thursday				//4
	Friday					//5
	Saturday				//6
)
*/
package ch2
import "fmt"

const boilingF = 212.0    //包一级的常量。同一目录下的所有go文件都可以访问该常量。但不能被外部包访问，因为开头字母是小写
type Celsius float64      //摄氏温度
type Fahrenheit float64   //华氏温度

const (
	AbsoluteZeroC Celsius = -273.15  //绝对零度
	FreezingC     Celsius = 0        //结冰点温度
	BoilingC      Celsius = 100      //沸腾温度
)

func Temperature() {
	/*
	需要注意的是。算术运算符，以及比较运算符可以用来比较一个命名类型的变量和另一个有相同类型的变量，或者
	有着相同底层类型的未命名类型的值之间做比较。例如，(f-i)会报错，但(f-32)不会，f - Celsius(i)也不会。
	*/
	i := 32
	var f Celsius = 100	  		//变量1
	var c = (f - Celsius(i)) * 5 / 9    //变量2
	fmt.Printf("boiling point is: %g°F or %g°C\n",f,c)
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

