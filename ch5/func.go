package ch5

import "fmt"

/*
参数数量可变的函数被称为可变参数函数。在声明可变参数函数时，需要在最后一个参数类型之前加上省略符号“...”
在函数体中，vals被看作是类型为[]int的切片。需要注意的是，可变参数只能作为函数参数列表中的最后一个参数。
*/
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func TestSum() {
	fmt.Println(sum())           // "0"
	fmt.Println(sum(3))          // "3"
	fmt.Println(sum(1, 2, 3, 4)) // "10"

	//当原始参数已经是一个切片时，在传递给函数的可变类型时，需要在最后一个参数后加上省略符
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))
}

func double(x int) (result int) {
	//匿名函数捕获的是调用他的函数体中的变量的地址，在某些场合中需要格外注意。
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func Parse(input string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()

	return err
}
