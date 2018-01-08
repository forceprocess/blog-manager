package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	. "utils"

	_ "github.com/go-sql-driver/mysql"
)

// SystemUser 用户实体
type SystemUser struct {
	Id            int    `json:"id"`
	LoginName     string `json:"loginName"`
	Name          string `json:"name"`
	IsAdmin       int    `json:"isAdmin"`
	Password      string `json:"password"`
	Status        int    `json:"status"`
	LastLoginTime string `json:"lastLoginTime"`
	Creator       string `json:"creator"`
	Updator       string `josn:"updator"`
	CreateTime    string `json:"createTime"`
	UpdateTime    string `json:"updateTime"`
	Phone         string `json:"phone"`
	Sex           int    `json:"sex"`
	Email         string `json:"email"`
}

// Login 登录页面
func Login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(Dir + "view/login.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}

// TOLogin 用于登录请求
func ToLogin(writer http.ResponseWriter, request *http.Request) {

	result := Result{}
	request.ParseForm()
	var buf []byte

	if request.Method == "POST" {

		username := request.PostFormValue("username")
		password := request.PostFormValue("password")

		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8")
		if err != nil {
			result.Message = "登录超时！"
			buf, _ = json.Marshal(result)
			writer.Header().Set("Content-Type", "application/json")
			writer.Write(buf)
			return
		}
		defer db.Close()

		sql := "select * from system_user where loginname=? and `password`=?"
		row, err := DB.Query(sql, username, password)
		if err != nil {
			result.Message = "登录失败！"
		}
		defer row.Close()

		if row.Next() {

			var systemUser SystemUser
			row.Scan(&systemUser.Id, &systemUser.LoginName, &systemUser.Name, &systemUser.IsAdmin, &systemUser.Password, &systemUser.Status, &systemUser.LastLoginTime, &systemUser.Creator, &systemUser.Updator, &systemUser.CreateTime, &systemUser.UpdateTime, &systemUser.Phone, &systemUser.Sex, &systemUser.Email)
			fmt.Println(systemUser)

			result.Success = true
			result.Message = "登录成功！"
		} else {
			result.Message = "用户名或密码错误！"
		}

	} else {
		result.Message = "登录方式错误！"
	}
	buf, _ = json.Marshal(result)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(buf)

	// if request.Method == "GET" {
	// 	err := request.ParseForm()
	// 	if err!=nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	if len(request.Form["username"]) > 0 {
	// 		fmt.Println(request.Form["username"][0])
	// 		fmt.Println(request.FormValue("username"))
	// 		fmt.Println(request.FormValue("password"))
	// 	}

	// } else {

	// 	err := request.ParseForm()
	// 	if err!=nil {
	// 		fmt.Println(err)
	// 		data := make(map[string]string)
	// 		data["msg"] = "登录失败！"
	// 		result.Data = data
	// 		buf,_ := json.Marshal(result)
	// 		writer.Header().Set("Content-Type","application/json")
	// 		writer.Write(buf)
	// 		return
	// 	}

	// 	//request.PostForm["username"] 数组
	// 	fmt.Println("username:", request.PostFormValue("username"))
	// 	fmt.Println("password:", request.PostFormValue("password"))

	// 	username := request.PostFormValue("username")
	// 	password := request.PostFormValue("password")

	// 	if username == "" || password == "" {
	// 		data := make(map[string]string)
	// 		data["msg"] = "登录失败！"
	// 		result.Data = data
	// 		buf,_ := json.Marshal(result)
	// 		writer.Header().Set("Content-Type","application/json")
	// 		writer.Write(buf)
	// 	} else {

	// 		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8")

	// 		if err != nil {
	// 			return
	// 		}
	// 		defer db.Close()

	// 		sql := "select * from system_user where loginname=? and `password`=?"
	// 		row,err := db.Query(sql, username, password)
	// 		if err!=nil {
	// 			data := make(map[string]string)
	// 			data["msg"] = "登录失败！"
	// 			result.Data = data
	// 		}
	// 		defer row.Close()

	// 		if row.Next() {

	// 			var systemUser SystemUser

	// 			cols, _ := row.Columns()
	// 			fmt.Println(cols)

	// 			row.Scan(&systemUser.Id, &systemUser.LoginName, &systemUser.Name, &systemUser.IsAdmin, &systemUser.Password, &systemUser.Status, &systemUser.LastLoginTime, &systemUser.Creator, &systemUser.Updator, &systemUser.CreateTime, &systemUser.UpdateTime, &systemUser.Phone, &systemUser.Sex, &systemUser.Email)
	// 			fmt.Println(systemUser)

	// 			data := make(map[string]string)
	// 			data["msg"] = "登录成功！"
	// 			result.Success = true
	// 			result.Data = data

	// 		} else {

	// 			data := make(map[string]string)
	// 			data["msg"] = "登录失败！"
	// 			result.Data = data
	// 		}

	// 	}
	// }
}
