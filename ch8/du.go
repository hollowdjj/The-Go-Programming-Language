package ch8

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)


var sema = make(chan struct{},20) //定义一个信号量，避免一次性打开太多文件(channel的阻塞性质)
var done = make(chan struct{})    //广播停止消息的channel

func cancelled() bool {
	select {
	case <- done:
		return true
	default:
		return false
	}
}

//以os.FileInfo切片的形式返回目录dirname下的所有内容
func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}

	defer func(){<-sema}()
	//ioutile.ReadDir(dirname string)以os.FileInfo切片的形式返回目录dirname下的所有内容
	entries,err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr,"du: %v\n",err)
		return nil
	}

	return entries
}

//将dir目录下的所有文件的文件大小写到fileSizes channel中
func walkDir(dir string,fileSizes chan<- int64,n *sync.WaitGroup) {
	defer n.Done()
	if cancelled() {
		return
	}

	for _,entry:= range dirents(dir) {
		//如果entry也是一个目录，那么就递归
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir,entry.Name())
			go walkDir(subdir,fileSizes,n)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var verbose = flag.Bool("v",false,"show verbose progress messages")
func Du() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	/*
	后台goroutine，用于找到目录下的所有文件。每一个walkDir调用都是在一个新的goroutine中。
	利用sync.WaitGroup进行计数，当所有文件都统计完毕后，关闭fileSizes
	*/
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _,root := range roots {
		n.Add(1)
		go walkDir(root,fileSizes,&n)
	}
	go func() {
		n.Wait()   		 //一直阻塞直到waitGroup的计时器为0
		close(fileSizes)
	}()

	//主goroutine
	var timer *time.Timer
	if *verbose {
		timer = time.NewTimer(100 * time.Microsecond)
	}

	var nfiles,nbytes int64
loop:
	for {
		select {
		case <- done:
			for _ = range fileSizes{} //清空fileSizes，否则walkDir会阻塞在：fileSizes <- entry.Size()语句
		case <-timer.C:
			printDiskUseage(nfiles,nbytes)
		case size,ok :=<- fileSizes:
			if !ok {
				break loop     //直接退出最外层的循环
			}
			nfiles++
			nbytes += size
		}
	}
	printDiskUseage(nfiles,nbytes)
}

func printDiskUseage(nfiles,nbytes int64) {
	fmt.Printf("%d files，%.1f GB\n",nfiles,float64(nbytes) / 1e9)
}