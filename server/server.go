/*
Webアプリケーション練習用
*/
package main

import (
	"log"
	"net/http"

	"github.com/ponta0929/goweb/handler"
)

//main()
func main() {
	//http.HandleFunc("/", handler.SampleHandler)
	http.HandleFunc("/view/", handler.ViewHandler)
	http.HandleFunc("/edit/", handler.EditHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
