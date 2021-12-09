package ch9_Variable_In_Concurency

import (
	"fmt"
	"sync"
)

/*
go中提倡的一种避免竞争的方法为：不要使用共享数据来通信，而是使用通信来共享数据。
实际上就是使用了一个容量为1的channel来保证最多只有一个goroutine在同一时刻访问
一个共享变量，充当的是一个互斥量的作用。
1. channel 2. sync.Mutex 3. sync.RWMutex 4. sync.Once
sync.Once和C++中的call_once一样，都是保证某个可能会被并发执行的函数中的某一语句永远只会执行一次
*/

/*
没有考虑并发读写的bank版本
*/
var balance int //银行总的余额

//Deposit 模拟用户存款
func Deposit(amount int) {
	balance += amount
}

//Balance 返回当前银行总的余额
func Balance() int {
	return balance
}

func Bank() {
	/*
		这里模拟了两个用户并发的进行存款操作。由于Deposit函数以及Balance函数对变量balance的读写均不是
		并发安全的，故这个函数是错误的。
	*/
	go func() {
		Deposit(100)
	}()

	go func() {
		Deposit(200)
	}()
}

/*
上面这个版本的Bank错误的原因在于，有多个goroutine并发去读写一个共享变量。
因此，一个改进方案是，使用一个monitor goroutine。共享变量的读写都通过这个
monitor goroutine完成。在这个版本的实现中，利用的就是非缓存channel阻塞性质。
*/
var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan bool)

//Deposit1 通过向deposits channel发生数据完成存款
func Deposit1(amount int) {
	deposits <- amount
}

//Balance1 通过向balances channel发生数据完成存款
func Balance1() int {
	return <-balances
}

//Withdraw 用户取款
func Withdraw(amount int) bool {
	Deposit1(-amount)
	return <-withdraw
}

func teller() {
	var totalBalance int
	for {
		select {
		case amount := <-deposits:
			//用户存款，直接相加
			if amount > 0 {
				totalBalance += amount
			} else {
				//用户取款，则需判断余额还够不够
				if totalBalance < -amount {
					withdraw <- false //余额不足
				} else {
					withdraw <- true
					totalBalance += amount //余额足
				}
			}
		case balances <- totalBalance:
			//Do nothing
		}
	}

}

func Bank1() {
	go teller()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Bob Deposit 100$")
		Deposit1(100)
		fmt.Printf("Bob——Current balance: %d\r\n", Balance1())
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Alice Deposit 200$")
		Deposit1(200)
		fmt.Printf("Alice——Current balance: %d\r\n", Balance1())
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("Emma withdraw 100$ %t\r\n", Withdraw(100))
		fmt.Printf("Emma: Current balance: %d\r\n", Balance1())
	}()

	wg.Wait()
}
