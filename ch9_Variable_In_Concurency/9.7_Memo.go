package ch9_Variable_In_Concurency

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

/*
设计一个缓存函数，也就是说我们需要缓存函数的结果，这样在对函数进行调用的时候，我们就只需要一次计算，之后
只要返回计算的结果就可以了。cache需要支持并发、不重复且无阻塞
*/

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	//ReadAll函数会返回一个[]byte和error，而空接口可以传任何值
	return ioutil.ReadAll(resp.Body)
}

//Func 要缓存结果的函数的类型
type Func func(key string) (interface{}, error)

//result 函数结果
type result struct {
	value interface{}
	err   error
}

//Memo 每一个Memo实例都会记录需要缓存的函数f以及不同的key值计算得到的结果
type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.Mutex
}

//entry
type entry struct {
	res   result
	ready chan struct{} //closed when res is ready
}

//New 返回一个缓存对象
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

//Get 访问函数结果
func (memo *Memo) Get(key string) (interface{}, error) {
	//查询cache时需要互斥访问
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		//没有条目则创建一个条目，随后解锁
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		//计算函数值。计算完后，关闭channel，相当于传播计算完毕的信号
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		//如果条目已经创建了，那么就等待那一个goroutine计算完成
		memo.mu.Unlock()
		<-e.ready //在调用f的那个goroutine没有填充完值之前(调用close函数)，试图读取该条目值的goroutine都会阻塞在这里。
	}
	return e.res.value, e.res.err
}

func inComingURLs() []string {
	return []string{"https://www.baidu.com/index.html", "https://www.4399.com/index.html"}
}

func TestMemo() {
	m := New(httpGetBody)
	for _, url := range inComingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s,%s,%d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}
