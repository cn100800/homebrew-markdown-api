package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	//"os"
)

var (
	//method string
	u string
)

func main() {
	// fmt.Println("enter method :")
	// fmt.Scanln(&method)
	// fmt.Println("method is :", method)
	fmt.Println("enter url :")
	fmt.Scanln(&u)
	fmt.Println("url is :", u)
	res, err := url.Parse(u)
	if err != nil {
		fmt.Println("url parse faild .")
	}
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.RawQuery)
	query, _ := url.ParseQuery(res.RawQuery)
	for k, v := range query {
		fmt.Println(k, ":", v)
	}
	hearder := "| 参数 | 备注 |\n|--- |---|\n"
	for k, v := range query {
		hearder = hearder + "|" + k + "|" + v[0] + "|\n"
	}

	//os.OpenFile("test.md", os.O_APPEND|os.O_CREATE|os.O_RDWR, 777)
	ioutil.WriteFile("test.md", []byte(hearder), 0777)
}
