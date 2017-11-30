package main

import "fmt"
import "net/url"
//import "strings"

func main() {
	s := "http://www.baidu.com"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u)
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.User)
}
