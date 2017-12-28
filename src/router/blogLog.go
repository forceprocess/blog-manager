package router

import (
	"net/http"
	"fmt"
	"html/template"
)

// BlogLog 点击
func BlogLog(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles(Dir + "view/blog-log.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}