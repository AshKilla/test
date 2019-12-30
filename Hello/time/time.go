package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() //获取当前时间
	fmt.Println("当前时间为：", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	mintue := now.Minute()
	second := now.Second()

	fmt.Printf("year:%v,month:%v,day:%v,hour:%v,mintue:%v,second:%v", year, month, day, hour, mintue, second)
}
