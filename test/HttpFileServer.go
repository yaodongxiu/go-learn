package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	fmt.Println("http file server start...")
	http.ListenAndServe(":9601", nil)
}
