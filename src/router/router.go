package router

import (
	"fmt"
	"net/http"
	"html/template"
)

func init() {

	fmt.Println("------路由开始初始化")

	//静态资源服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//登录页
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("view/login.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(writer, nil)
	})

	//管理后台首页
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles("view/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(writer, nil)
	})

	fmt.Println("------路由初始化完成")
}
