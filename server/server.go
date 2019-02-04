package main

import (
	"fmt"
	"io/ioutil"
)

/*
Page
ページを表す構造体
*/
type Page struct {
	Title string
	Body  []byte
}

//save() ページ情報をタイトル名でファイルに保存する
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//load() 保存したページを読み込む
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

//main()
func main() {
	p1 := &Page{Title: "サンプルページ", Body: []byte("これはサンプルページです")}
	p1.save()
	p2, _ := loadPage("サンプルページ")
	fmt.Println(string(p2.Body))
}
