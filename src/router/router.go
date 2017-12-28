package router

import (
	"strings"
	"os"
	"fmt"
	"net/http"
)

// Dir 项目路径
var Dir string

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
	// 全局Dir
	Dir = dir

	fmt.Println("------路由开始初始化")

	//静态资源服务
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir + "static"))))

	//登录页
	http.HandleFunc("/login", Login)

	//首页
	http.HandleFunc("/index", Index)

	//登录
	http.HandleFunc("/toLogin", ToLogin)

	// 查询菜单
	http.HandleFunc("/queryMenuList",QueryMenuList)

	// 文章管理
	http.HandleFunc("/article", Article)

	// 用户管理
	http.HandleFunc("/user",User)

	// 访问统计
	http.HandleFunc("/access",Access)

	//点击统计
	http.HandleFunc("/click", Click)
	
	//博客日志
	http.HandleFunc("/blogLog", BlogLog)

	//后台日志
	http.HandleFunc("/managerLog", ManagerLog)

	fmt.Println("------路由初始化完成")
}
