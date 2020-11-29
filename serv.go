package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
)

const retStr = "<!doctype html>\r\n<html>\r\n  <head>\r\n    <meta charset=\"utf-8\">\r\n    <title></title>\r\n  </head>\r\n  <body>\r\n    <p>it works!</p>\r\n  </body>\r\n</html>\r\n"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	// fmt.Println("Hello world!")
	fmt.Println("running")
	fmt.Println(retStr)

	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, retStr)
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello world!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	// s := &http.Server{
	// 	Addr:           ":8080",
	// 	Handler:        defaultHandler,
	// 	ReadTimeout:    5 * time.Second,
	// 	WriteTimeout:   5 * time.Second,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	log.Fatal(http.ListenAndServe(":8080", nil))
}
