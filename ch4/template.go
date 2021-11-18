package ch4

import (
	"html/template"
	"log"
	"os"
	"time"
)

/*
一个模板是一个字符串或文件，里面包含了一个或多个由花括号包含的{{action}}对象。大部分的字符串只是按字面值打印，
但是对于action部分将触发特定的行为。action中的"."表示当前值，即最初被初始化为调用模板时的参数。{{.TotalCount}}
表示展开结构体中的TotalCount成员。{{range .Items}}和{{end}}对应一个循环action。在一个action中，|操作符表示将
前一个表达式的结果作为后一个函数的输入
*/
const templ = `{{.TotalCount}} issues:            
{{range .Items}}--------------------------------
Number: {{.Number}}
User:	{{.User.Login}}
Title: 	{{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours()) / 24
}

func TestTemplate() {
	//首先需要分析模板并转为内部表示。这里即为创建模板并注册daysAgo函数，并用Must函数进行编译期的模板检查
	report := template.Must(template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo}).Parse(templ))

	//获取json数据并解码
	res, err := SearchIssue(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	//利用text模板打印json数据
	if err := report.Execute(os.Stdout, res); err != nil {
		log.Fatal(err)
	}
}
