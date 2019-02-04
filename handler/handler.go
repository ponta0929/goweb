// Package handler 練習用のハンドラーパッケージ
package handler

import (
	"fmt"
	"net/http"
)

//SampleHandler func(w http.ResponseWriter, r *http.Request)
func SampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
