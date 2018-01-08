package utils

import (
	"os"
	"fmt"
	"strings"
)

var (
	CtxPath string
)

func init() {
	CtxPath = ContextPath()
}

// ContextPath 获取上下文路径
func ContextPath() string {

	dir, err := os.Getwd()
	dir = dir + string(os.PathSeparator)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if strings.Index(dir, "src") > -1 {
		dir = dir[0:strings.Index(dir, "src")]
	}
	CtxPath = dir
	return dir
}
