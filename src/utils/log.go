package utils

import (
	"log"
	"os"
	"io"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {

	errFile, err := os.OpenFile(logFilePath(), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	Info = log.New(os.Stdout, "Info:", log.LstdFlags|log.Lshortfile)
	Warning = log.New(os.Stdout, "Warning:", log.LstdFlags|log.Lshortfile)
	Error = log.New(io.MultiWriter(os.Stderr, errFile), "Error:", log.LstdFlags|log.Lshortfile)
}

//logFilePath 日志文件路径
func logFilePath() string {
	ctxPath := ContextPath()
	return ctxPath + "logs/blog-manager.log"
}
