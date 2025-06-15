package main

import (
	"fmt"
	"github.com/xaluk08/gift-service-new/internal/hello"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	NewHelloHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("Listening on port 8081")
	server.ListenAndServe()

}
