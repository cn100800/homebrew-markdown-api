package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"time"
)

var (
	method string
	u      string
)

func main() {
	//p := fmt.Println
	fmt.Println("enter method :")
	fmt.Scanln(&method)
	fmt.Println("method is :", method)
	fmt.Println("enter url :")
	fmt.Scanln(&u)
	//fmt.Println("url is :", u)
	res, err := url.Parse(u)
	if err != nil {
		fmt.Println("url parse faild .")
	}
	// fmt.Println(res.Host)
	// fmt.Println(res.Path)
	// fmt.Println(res.RawQuery)
	query, _ := url.ParseQuery(res.RawQuery)
	hearder := "|method|path|\n|---|---|\n|" + method + "|" + res.Path + "|\n#query\n| 参数 | 备注 |\n|--- |---|\n"
	for k, v := range query {
		hearder = hearder + "|" + k + "|" + v[0] + "|\n"
	}
	hearder = hearder + "```json\n```"
	//os.OpenFile("test.md", os.O_APPEND|os.O_CREATE|os.O_RDWR, 777)
	t := time.Now()
	file_name := []byte(t.Format(time.RFC3339))[0:10]
	ioutil.WriteFile("/Users/cn/Desktop/"+string(file_name)+".md", []byte(hearder), 0777)
}
