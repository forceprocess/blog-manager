package router

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"utils"

	_ "github.com/go-sql-driver/mysql"
)

// Menu Modal
type Menu struct {
	Id          string `json:"id"`
	Pid         string `json:"pid"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Icon        string `json:"icon"`
	Status      string `json:"status"`
	Creator     string `json:"creator"`
	Updator     string `json:"updator"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
	Description string `json:"description"`
	Children    []Menu `json:"children"`
}

// QueryMenuList 获取菜单
func QueryMenuList(w http.ResponseWriter, r *http.Request) {

	result := utils.Result{}
	var buf []byte

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err != nil {
		result.Message = "获取菜单超时！"
		buf, _ = json.Marshal(result)
		w.Header().Set("Content-Type", "application/json")
		w.Write(buf)
		return
	}
	defer db.Close()

	sql := "select id,pid,name,path,icon from system_menu where status=0 and pid=?"
	row, err := db.Query(sql, 0)
	if err != nil {
		result.Message = "获取菜单失败！"
	}
	defer row.Close()

	menus := make([]Menu, 0)
	for row.Next() {

		var menu Menu
		row.Scan(&menu.Id, &menu.Pid, &menu.Name, &menu.Path, &menu.Icon)

		children, err := db.Query(sql,menu.Id)
		if err!=nil {
			fmt.Println(err)
			return
		}
		defer children.Close()

		childrens := make([]Menu, 0)
		for children.Next() {
			var me Menu
			children.Scan(&me.Id, &me.Pid, &me.Name, &me.Path, &me.Icon)
			childrens = append(childrens,me)
		}
		menu.Children = childrens

		menus = append(menus, menu)
		result.Success = true
	}
	result.Data = menus
	buf, _ = json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
