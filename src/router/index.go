package router

import (
	"net/http"
	"fmt"
	"html/template"
)

// Index 访问统计
func Index(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles(Dir + "view/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}