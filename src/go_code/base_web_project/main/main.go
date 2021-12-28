package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"project2/src/go_code/base_web_project/web_server"
)

func main() {
	//NetHttpTest()
	HttprouterTest()
}

func HttprouterTest() {

	router := httprouter.New()

	router.GET("/getTest", web_server.GetTest)
	router.GET("/getTest/:name/:age", web_server.GetUrlParamTest)

	http.ListenAndServe(":8888", router)

}

// go本身带的net/http各接口的测试
func NetHttpTest() {
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
