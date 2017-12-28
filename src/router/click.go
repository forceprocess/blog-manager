package router

import (
	"net/http"
	"fmt"
	"html/template"
)

// Click 点击
func Click(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles(Dir + "view/click.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}