package ch7

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32   				 		//定义一个美元类型
type database map[string]dollars	 		//数据库，即物品名称与其价格的映射

//需要正对dollars类型实现一个String方法以实现打印值
func (d dollars) String() string{
	return fmt.Sprintf("$%.2f",d)    //小数点后只输出两位数字
}

//db类型实现了ServeHttp方法，即实现了http.Hander接口，这个版本的实现没有考虑URL
//func (db database) ServeHTTP(w http.ResponseWriter,req *http.Request){
//	for item,price := range db {
//		fmt.Fprintf(w,"%s: %s\n",item,price)
//	}
//}

//考虑了URL的第二个实现版本，但更为理想的做法是将每个case拆分成一个单独的函数
func (db database) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item,price := range db {
			fmt.Fprintf(w,"%s, %s\n",item,price)
		}
	case "/price":         //查找指定item的价格URL示例: /price?item=socks
		item := req.URL.Query().Get("item")
		price,ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)  //发送404的相关header
			fmt.Fprintf(w,"no such item: %q\n",item)
			return
		}
		fmt.Fprintf(w,"%s\n",price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w,"no such page: %q\n",req.URL)
	}
}

/*
通过对database类型编写符合ServeHTTP接口list和price函数实现不同URL的响应处理，即：
http.HandleFunc("/list",db.list)
http.HandleFunc("/price",db.price)
*/
func (db database) list(w http.ResponseWriter,req *http.Request){
	for item,price := range db {
		fmt.Fprintf(w,"%s: %s\n",item,price)
	}
}

func (db database) price(w http.ResponseWriter,req *http.Request){     //usage: /price?item=socks
	item := req.URL.Query().Get("item")
	price,ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)  		//发送404的相关header
		fmt.Fprintf(w,"no such item: %q\n",item)
		return
	}
	fmt.Fprintf(w,"%s\n",price)
}

func (db database) create(w http.ResponseWriter,req *http.Request){
	item := req.URL.Query().Get("item")

	_,ok := db[item]
	if ok{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"item: %s already exists\n",item)
		return
	}

	price,_ := strconv.ParseFloat(req.URL.Query().Get("price"),32)
	db[item] = dollars(price)
	fmt.Fprintf(w,"Add item %q success! Current list:\n",item)
	db.list(w,req)
} //usage: /create?item=socks&price=6

func (db database) update(w http.ResponseWriter,req *http.Request){
	item := req.URL.Query().Get("item")

	oldPrice,ok :=db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"item: %s does not exist\n",item)
		return
	}

	newPrice,_ := strconv.ParseFloat(req.URL.Query().Get("price"),32)
	db[item] = dollars(newPrice)
	fmt.Fprintf(w,"Update success!\nBefore: %s: %s\nNow: %s: %s\n",item,oldPrice,item, db[item])
} //usage: /update?item=socks&price=6

func (db database) delete(w http.ResponseWriter,req *http.Request) {
	item := req.URL.Query().Get("item")

	_,ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w,"item: %s does not exist\n",item)
		return
	}

	delete(db,item)
	fmt.Fprintf(w,"Delete item %q success! Current list:\n",item)
	db.list(w,req)
}

func Http() {
	db := database{"socks" : 1,"pants": 5}
	/*
	http.ListenAndServer函数接受两个参数。第一个参数为监听的地址，第二个参数是一个接口，该接口要求
	类型实现ServeHTTP(ResponseWriter, *Request)方法，该方法用于响应客户端的http请求
	*/
	http.HandleFunc("/list",db.list)   //这些注册的handler是在一个新的gocoroutine上被调用
	http.HandleFunc("/price",db.price)
	http.HandleFunc("/create",db.create)
	http.HandleFunc("/update",db.update)
	http.HandleFunc("/delete",db.delete)

 	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}


