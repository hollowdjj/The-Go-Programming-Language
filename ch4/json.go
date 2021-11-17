package ch4

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

/*
JSON是对JavaScript中各种类型的值——字符串、数字、布尔值和对象
*/

type Movie struct {
	Title  string
	Year   int  `json:"released"`        //Go中的Year成员对应到JSON中的released对象(主要是因为JSON与Go的编码命名规范不同)
	Color  bool `json:"color,omitempty"` //表示Go中的结构体成员为空或者零值时不生成JSON对象
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

	//解码JSON数据并且可以有选择性的解码JSON中感兴趣的成员。例如Title
	var titles []struct{ Title string }
	if err = json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshal failed: %s", err)
	}
	fmt.Println(titles)
}

/*
访问Github提供的JSON接口
*/

const IssuesURL = "https://api.github.com/search/issues"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

func SearchIssue(terms []string) (*IssueSearchResult, error) {
	//url.QueryEscape用于将查询中的特殊字符(类似?和&)进行转义操作
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssueSearchResult
	//NewDecoder是一个基于流式的解码器
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func Practice410() {
	res, err := SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now()
	lessThanOneMonth, lessThanOneYear, moreThanOneYear := []*Issue{}, []*Issue{}, []*Issue{}

	for _, v := range res.Items {
		if v.CreatedAt.Month() == now.Month() {
			lessThanOneMonth = append(lessThanOneMonth, v)
		}
		if v.CreatedAt.Year() == now.Year() {
			lessThanOneYear = append(lessThanOneYear, v)
		} else if v.CreatedAt.Year() < now.Year() {
			moreThanOneYear = append(moreThanOneYear, v)
		}
	}

	fmt.Printf("There are %d issues created less than 1 month ago:\n", len(lessThanOneMonth))
	for _, item := range lessThanOneMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("There are %d issues created less than 1 Year ago:\n", len(lessThanOneYear))
	for _, item := range lessThanOneYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
	fmt.Println("---------------------------------------------------------------")
	fmt.Printf("There are %d issues created more than 1 Year ago:\n", len(moreThanOneYear))
	for _, item := range moreThanOneYear {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func TestGithubWebAPIS() {
	res, err := SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", res.TotalCount)
	for _, item := range res.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

/*
练习4.11：编写一个工具，允许用户在命名行创建、读取、更新和关闭Github上的issue，当必要的时候自动打开用户默认
的编辑器用于输入文本信息。
*/
var create = flag.String("c", "", "create an issue, usage: -c [owner] [repo] [title] [content]")
var delete = flag.String("d", "", "delete an issue, usage: -d [owner] [repo]")
var close = flag.String("s", "", "shut down an issue,usage: -d [owner] [repo] [title]")

func Practice411() {

}
