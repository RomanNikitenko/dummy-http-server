package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	path := "/"
	if len(os.Args) >= 2 {
		port = ":" + os.Args[1]
	}
	if len(os.Args) >= 3 {
		path = os.Args[2]
	}

	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello World!"))
    })

    http.HandleFunc("/test1", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Test resource "))
    })

    http.HandleFunc("/test2", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Another resource "))
	})
	
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
	fmt.Println("Done ...")
}
