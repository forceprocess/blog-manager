package router

import (
	"net/http"
	"fmt"
	"html/template"
)

// ManagerLog 点击
func ManagerLog(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles(Dir + "view/manager-log.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}