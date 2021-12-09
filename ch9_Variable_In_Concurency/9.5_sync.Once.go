package ch9_Variable_In_Concurency

import "sync"

/*
sync.Once。即C++中的std::call_once，保证函数只会被调用一次
*/

var icons map[string]string
var loadIconsOnce sync.Once

func loadIcons() {
	//假设这是一个比较消耗资源的初始化操作
	icons = map[string]string{
		"spades.png":   "spades",
		"hearts.png":   "hearts",
		"diamonds.png": "diamonds",
		"clubs.png":    "clubs",
	}
}

func Icon(name string) string {
	//每一次对Do的调用都会锁定一个Mutex，并会检查一个boolean变量。
	//在首次调用Do时，boolean变量的值为false，本次调用会调用loadIcons函数并
	//将其boolean变量设置为true。后续Do函数的调用什么都不会做。
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}
