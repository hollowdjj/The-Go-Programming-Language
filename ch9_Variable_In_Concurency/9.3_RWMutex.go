package ch9_Variable_In_Concurency

import "sync"

/*
读写锁
对于一个共享变量，如果只是读取值而不是修改值的话，完全没有必要上锁。也就是说，可以允许同时读，但不能
允许同时写，或者同时读写。读写锁就提供了这样一个功能。
*/
var mu sync.RWMutex

func Balance4() int {
	mu.RLock() //注意，这里要使用读锁
	defer mu.RUnlock()
	return balance
}

func Deposit4(amount int) {
	mu.Lock() //这里是写锁
	defer mu.Lock()
	balance += amount
}
