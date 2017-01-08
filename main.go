package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/react", react)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	fmt.Println("listening... port" + port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
}

func react(res http.ResponseWriter, req *http.Request) {
	println("react")
	http.ServeFile(res, req, "./index.html")
}
