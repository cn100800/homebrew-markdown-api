package version

import (
	"fmt"
)

var (
	VERSION string
)

func init() {
	//fmt.Println(VERSION)
}

func GetVersion() {
	fmt.Println(VERSION)
}
