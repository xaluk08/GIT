package hello

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(w, "Hello ret")
}

type HelloHandler struct {
}

func NewHelloHandler(router *http.ServeMux) {
	handler := HelloHandler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (handler *HelloHandler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Hello handler")
	}
}
