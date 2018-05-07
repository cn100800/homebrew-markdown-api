package main

import (
	"flag"
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
	// -u -m -t
	f := flag.String("f", "", "file name")
	t := flag.String("t", "api title", "api title")
	m := flag.String("m", "request method", "request method")
	flag.StringVar(&u, "u", "request url", "request url")
	flag.Parse()
	// if *m == "" {
	// 	fmt.Println("enter method :")
	// 	fmt.Scanln(m)
	// 	fmt.Println("method is :", m)
	// }
	// if u == "" {
	// 	fmt.Println("enter url :")
	// 	fmt.Scanln(&u)
	// 	fmt.Println("url is :", u)
	// }
	res, err := url.Parse(u)
	if err != nil {
		fmt.Println("url parse faild .")
	}
	// fmt.Println(res.Host)
	// fmt.Println(res.Path)
	// fmt.Println(res.RawQuery)
	query, _ := url.ParseQuery(res.RawQuery)
	hearder := "#" + *t + "\n|method|path|\n|---|---|\n|" + *m + "|" + res.Path + "|\n#query\n| 参数 | 备注 |\n|--- |---|\n"
	for k, v := range query {
		hearder = hearder + "|" + k + "|" + v[0] + "|\n"
	}
	hearder = hearder + "#result\n" + "```json\n```"
	//os.OpenFile("test.md", os.O_APPEND|os.O_CREATE|os.O_RDWR, 777)
	t_now := time.Now()
	file_name := []byte(t_now.Format(time.RFC3339))[0:10]
	//fmt.Println(hearder)
	if *f == "" {
		ioutil.WriteFile("/Users/cn/Desktop/"+string(file_name)+".md", []byte(hearder), 0777)
	} else {
		ioutil.WriteFile("/Users/cn/apidoc/"+*f+".md", []byte(hearder), 0777)
	}
}
