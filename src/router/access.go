package router

import (
	"net/http"
	"fmt"
	"html/template"
)

// Access 访问统计
func Access(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles(Dir + "view/access.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}