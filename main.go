package main

import (
	"fmt"
	"mojito"
	"net/http"
)

// version 1:
// func main() {
// 	http.HandleFunc("/", indexHandler)
// 	http.HandleFunc("/hello", helloHandler)
// 	log.Fatal(http.ListenAndServe(":9999", nil))
// }

// func indexHandler(writer http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(writer, "URL.Path = %q\n", req.URL.Path)
// }

// func helloHandler(writer http.ResponseWriter, req *http.Request) {
// 	for k, v := range req.Header {
// 		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
// 	}
// }

// version 2:
// type Engine struct {
// }

// func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	switch request.URL.Path {
// 	case "/":
// 		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
// 	case "/hello":
// 		for k, v := range request.Header {
// 			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
// 		}
// 	default:
// 		fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL)
// 	}
// }

// func main() {
// 	engine := new(Engine)
// 	log.Fatal(http.ListenAndServe(":9999", engine))
// }

func main() {
	r := mojito.New()

	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})

	r.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range req.Head {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
