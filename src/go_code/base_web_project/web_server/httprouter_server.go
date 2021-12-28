package web_server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func GetTest(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	fmt.Fprintf(writer, "my name is %s, age is %d", "吴福强", 20)
}

func GetUrlParamTest(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	fmt.Fprintf(writer, "my name is %s, age is %s", params.ByName("name"), params.ByName("age"))
}
