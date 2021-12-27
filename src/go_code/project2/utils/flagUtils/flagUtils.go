package flagUtils

import (
	"flag"
	"fmt"
)

func ParseArgs() {
	fmt.Println("-------flag解析参数-------------")
	var host string
	var port int
	var username string
	var password string

	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "p", 3306, "端口，默认为3306")
	flag.StringVar(&username, "username", "", "用户名")
	flag.StringVar(&password, "pwd", "", "密码")
	flag.Parse()

	fmt.Printf("参数信息：host=%v,port=%v,username=%s,password=%s \n", host, port, username, password)
}
