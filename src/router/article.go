package router

import (
	"fmt"
	"html/template"
	"net/http"
)

// Article 文章管理页面
func Article(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles(Dir + "view/article.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}
