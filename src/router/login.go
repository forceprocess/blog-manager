package router

import (
	"net/http"
	"fmt"
	"io"
)

type User struct {
	 Username string
	 Password string
}

// TOLogin 用于登录请求
func ToLogin(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		err := request.ParseForm()
		if err!=nil {
			fmt.Println(err)
			return
		}

		if len(request.Form["username"]) > 0 {
			fmt.Println(request.Form["username"][0])
			fmt.Println(request.FormValue("username"))
			fmt.Println(request.FormValue("password"))
		}

	} else {
		err := request.ParseForm()
		if err!=nil {
			fmt.Println(err)
			return
		}

		//request.PostForm["username"] 数组
		fmt.Println("username:", request.PostFormValue("username"))
		fmt.Println("password:", request.PostFormValue("password"))
	}
	io.WriteString(writer,"{\"code\":1,\"result\":\"登录成功\"}")
}
