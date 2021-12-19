package main

import (
	"fmt"
	"net/http"
	"project2/src/go_code/base_web_project/web_server"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/getName", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("请求参数：", request)
		fmt.Fprint(writer, "I'm a server,my name is wufuqiang")
	})

	mux.HandleFunc("/getAge", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "I'm a server,my age is 19")
	})

	mux.HandleFunc("/getHeader", web_server.GetHttpHeader)

	mux.HandleFunc("/getParam", web_server.GetHttpHandler)

	http.ListenAndServe(":8888", mux)
}
