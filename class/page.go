/*
Package class Webアプリケーション練習構造体パッケージ
*/
package class

import (
	"io/ioutil"
)

/*
Page ページを表す構造体
*/
type Page struct {
	Title string
	Body  []byte
}

//save() ページ情報をタイトル名でファイルに保存する
func (p *Page) save() error {
	filename := "../static/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//LoadPage 保存したページを読み込む
func LoadPage(title string) (*Page, error) {
	filename := "../static/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
