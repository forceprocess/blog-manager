package utils

import (
	"testing"
	"fmt"
)

func TestContextPath(t *testing.T) {
	ctxPath := ContextPath()
	fmt.Println(ctxPath)
}

func TestQueryUser(t *testing.T) {
	sql := "select id from t_user"
	row, err := DB.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer row.Close()

	for row.Next()  {
		var id string
		row.Scan(&id)
		fmt.Println(id)
	}

}

func TestLog(t *testing.T) {
	Info.Println("Info")
	Error.Println("Error")
	Error.Fatal("0000")
	Warning.Printf("Warning：%s\n","Warning")
}