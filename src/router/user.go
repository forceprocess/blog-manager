package router

import (
	"net/http"
	"fmt"
	"html/template"
)

// User 用户管理
func User(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles(Dir + "view/user.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}