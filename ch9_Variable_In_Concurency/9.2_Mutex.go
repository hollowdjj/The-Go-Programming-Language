package ch9_Variable_In_Concurency

import "sync"

/*
使无缓存channel模拟信号量，实现共享变量的并发访问
*/
var sema = make(chan struct{})

func Deposit2(amount int) {
	sema <- struct{}{}
	balance += amount
	<-sema
}

func Balance2() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

/*
直接使用sync包中的互斥锁，实现共享变量的并发访问
*/
var m sync.Mutex

func Deposit3(amount int) {
	m.Lock()
	defer m.Unlock()
	balance += amount
}

func Balance3() int {
	m.Lock()
	defer m.Unlock()
	return balance
}

func Withdraw3(amount int) bool {
	m.Lock()
	defer m.Unlock()
	if balance < amount {
		return false
	}
	balance -= amount
	return true
}
