package main

import (
	"log"
	"os"
	"strings"
)

const (
	LogPath = "./logs/1.log"
)

func CreateFile(filePath string) (*os.File, error) {
	fps := strings.Split(filePath, "/")
	if len(fps) > 1 {
		dir := strings.Join(fps[:len(fps)-1], "/")
		os.MkdirAll(dir, os.ModePerm)
	}
	return os.Create(filePath)
}

func OpenCreateFile(filePath string) (*os.File, error) {
	// 按照所需读写权限创建文件
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			f, err = CreateFile(filePath)
		}
	}
	return f, err
}

func main() {
	f, err := OpenCreateFile(LogPath)
	if err != nil {
		log.Fatalln("create log file fatal path:", LogPath, "; err:", err)
		panic(err)
	}
	// 完成后延迟关闭
	defer f.Close()
	//设置日志输出到 f
	log.SetOutput(f)

	log.Println("main error")
}
