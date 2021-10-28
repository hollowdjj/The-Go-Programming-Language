package ch8

import (
	"fmt"
	"time"
)

func RunSpinner() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n",n,fibN)
}

func spinner(delay time.Duration) {
	for  {
		for _,r:= range `-\|/` {     //反引号为原生字符串
			fmt.Printf("\r%c",r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x - 1) + fib(x - 2)
}