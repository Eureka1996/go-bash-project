package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"go_base/src/go_code/base_web_project/web_server"
	"net/http"
)

func main() {
	//NetHttpTest()
	//HttprouterTest()
	GinTest2()
}

func GinTest2() {
	r := web_server.GetEngine()

	r.Run(":8080")
}

func GinTest() {
	fmt.Println("wufuqiang")
	engine := gin.Default()
	group := engine.Group("/v1")
	{
		group.GET("/users", func(c *gin.Context) {
			fmt.Println("哈哈哈")
			c.String(200, "wufuqiang")
		})
	}

	group2 := engine.Group("/v2", func(c *gin.Context) {
		fmt.Println("/v2中间件")
	})
	{
		group2.GET("/user", func(c *gin.Context) {
			c.String(200, "v2_user")
		})
	}

	engine.Run(":8080")
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
