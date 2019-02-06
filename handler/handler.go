// Package handler 練習用のハンドラーパッケージ
package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ponta0929/goweb/class"
)

//SampleHandler func(w http.ResponseWriter, r *http.Request)
func SampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

//ViewHandler func(w http.ResponseWriter, r *http.Request)
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	t, _ := template.ParseFiles("../static/view.html")
	renderTemplate(t, w, title)
}

//EditHandler func
func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	t, _ := template.ParseFiles("../static/edit.html")
	renderTemplate(t, w, title)
}

func renderTemplate(t *template.Template, w http.ResponseWriter, title string) {
	p, err := class.LoadPage(title)
	if err != nil {
		p = &class.Page{Title: title}
	}
	t.Execute(w, p)
}
