package web_server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHttpHandler(response http.ResponseWriter, request *http.Request) {
	GetHttpHeader(response, request)
	GetHttpParam(response, request)
	GetHttpBody(response, request)
}

func GetHttpHeader(response http.ResponseWriter, request *http.Request) {

	// req.Header中Header本质是：type Header map[string][]string
	header := request.Header

	var acc []string = header["Accept"]

	for _, n := range acc {
		fmt.Fprintln(response, "Accept内容：", n)
	}

	fmt.Fprintln(response, "你发送的请求地址是：", request.URL.Path)
	fmt.Fprintln(response, "你发送的请求地址后查询字符串是：", request.URL.RawQuery)
	method := request.Method
	fmt.Println("请求类型：", method)
	fmt.Fprintln(response, "请求类型：", method)
}

func GetHttpParam(response http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Fprintln(response, request.Form)

	name := request.FormValue("name")
	age := request.FormValue("age")

	fmt.Fprintln(response, "姓名：", name, "，年龄：", age)

}

func GetHttpBody(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	if method == "POST" {
		fmt.Println("Post请求，获取Body")
		allBody, _ := ioutil.ReadAll(request.Body)
		fmt.Println(string(allBody))
		fmt.Fprintln(response, "请求的Body：", string(allBody))
	}
}
