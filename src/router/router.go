package router

import (
	"strings"
	"os"
	"fmt"
	"net/http"
	"html/template"
)

func init() {

	//获取当前所在路径 (此方式只适合main在blog-manager文件夹中启动)
	dir,err := os.Getwd()
	dir = dir + string(os.PathSeparator)
	if err != nil {
		fmt.Println(err)
		return
	}
	if strings.Index(dir,"src") > -1 {
		dir = dir[0:strings.Index(dir,"src")]
	}

	fmt.Println("------路由开始初始化")

	//静态资源服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir + "static"))))

	//登录页
	http.HandleFunc("/login", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles(dir+"view/login.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(writer, nil)
	})

	//管理后台首页
	http.HandleFunc("/index", func(writer http.ResponseWriter, request *http.Request) {
		t, err := template.ParseFiles(dir+"view/index.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(writer, nil)
	})

	//登录
	http.HandleFunc("/toLogin", ToLogin)

	fmt.Println("------路由初始化完成")
}
