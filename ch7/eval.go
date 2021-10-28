package ch7

import (
	"fmt"
	"math"
)
//Env 将Var映射为其对应的值
type Env map[Var]float64

//Expr 表示Go中任意的表达式
type Expr interface {
	//Eval 表达式求值
	Eval(env Env) float64
}

//Var 表示变量,例如x,是所有表达式中最基础的一个，故必须要一个映射指明这个变量的数值
type Var string

func (v Var) Eval(env Env) float64 {
	return env[v]
}

//literal 表示一个数值常量，如3,1415
type literal float64

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

//一元操作符表达式，如-x +x
type unary struct {
	op rune   //+或-
	x Expr
}

//Eval unary实现Expr接口
func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}

	panic(fmt.Sprintf("unsupport unary operator: %q",u.op))
}

//二元操作符表达式，如x+y
type binary struct {
	op rune //+ - * /
	x,y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}

	panic(fmt.Sprintf("unsupport unary operator: %q",b.op))
}

//函数调用表达式，如sin(x)
type call struct {
	fn string   //函数，pow sin sqrt
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env),c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}

	panic(fmt.Sprintf("unsupport unary operator: %q",c.fn))
}

//func TestEval(t* testing.T) {
//	tests := []struct{
//		expr string      //用字符串表示的表达式
// 		env Env          //表达式expr中使用到的变量的数值映射
//		want string      //表达式的正确值
//	}{
//		{"sqrt(A / pi)",Env{"A": 87616,"pi": math.Pi},"167"},
//		{"pow(x,3) + pow(y,3)",Env{"x": 2,"y": 1},"9"},
//		{"5 / 9 * (F - 32)",Env{"F": -40},"-40"},
//	}
//
//	var prevExpr string
//	for _,test := range tests {
//
//	}
//}
