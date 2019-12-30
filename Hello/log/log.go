package main

import (
	"fmt"
	"log"
	"os"
)

//日志标准库的使用

func main() {

	logFile, err := os.OpenFile("./chenjai.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open failed:err:", err)
		return
	}

	log.SetOutput(logFile)

	log.SetPrefix("[chenjia]") //设置前缀
	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")

	log.SetFlags(log.Llongfile | log.Lshortfile | log.Ldate) //设置flag
	log.Panicln("这是一条会触发panic的日志。")
}
