package ch4

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
JSON是对JavaScript中各种类型的值——字符串、数字、布尔值和对象
*/

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func TestJson() {
	//这样的数据结构特别适合JSON格式，将一个Go语言中类似movies的结构体slice转为JSON的过程叫编组
	movies := []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	}
	//编组marshaling可通过调用json.Marshal函数完成
	//这里的MarshalIndent函数将产生整齐缩进的输出，prefix表示每一行输出的前缀和每一层级的缩进
	data, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}
