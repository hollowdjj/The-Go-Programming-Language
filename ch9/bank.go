package ch9

import "sync"

/*
go中提倡的一种避免竞争的方法为：不要使用共享数据来通信，而是使用通信来共享数据。
实际上就是使用了一个容量为1的channel来保证最多只有一个goroutine在同一时刻访问
一个共享变量，充当的是一个互斥量的作用。
1. channel 2. sync.Mutex 3. sync.RWMutex 4. sync.Once
sync.Once和C++中的call_once一样，都是保证某个可能会被并发执行的函数中的某一语句永远只会执行一次
*/

var deposits = make(chan int)
var balances = make(chan int)
var (
	mu sync.Mutex
	murw sync.RWMutex    //读写锁
	balance int
)

//存款
func Deposits(amount int) {
	deposits <- amount
}

func Deposits1(amount int) {
	mu.Lock()
	balance += amount
	mu.Unlock()
}

//余额
func Balance() int {
	return <-balances
}

func Balance1() int {
	murw.RLock()                 //获取一个读锁
	defer murw.RUnlock()
	return balance
}

//取款
func Withdraw(amount int) bool {
	mu.Lock()
	defer  mu.Unlock()
	balance -= amount
	if balance < 0 {
		balance += amount
		return false
	}
	return  true
}

//在go中这种提供通过channel对一个指定变量进行访问的goroutine被称为monitor goroutine
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits: //只要调用了Deposits函数，这个case就会激活
			balance += amount
		case balances <- balance: //一开始时，会写一次，然后阻塞(非缓存channel)。调用Balance函数取出值后，激活
		}
	}
}

func init() {
	go teller()
}