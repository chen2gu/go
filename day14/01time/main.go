package main

import (
	"fmt"
	"time"
)

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分
	fmt.Println(now.Format("2006-01-02 15:04:05")) //2019-04-24 18:37:59
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func timeDemo() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Format("2006-01-02 15:04"))
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	//输出格式可以自定义
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)                                                  //将时间戳转为时间格式
	fmt.Println(timeObj)                                                                //2019-04-24 18:30:12 +0800 CST
	year := timeObj.Year()                                                              //年
	month := timeObj.Month()                                                            //月
	day := timeObj.Day()                                                                //日
	hour := timeObj.Hour()                                                              //小时
	minute := timeObj.Minute()                                                          //分钟
	second := timeObj.Second()                                                          //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second) //2019-04-24 18:30:12
}

func main() {
	fmt.Println("timeDemo.------->")
	timeDemo()
	//fmt.Println("timestampDemo2.------->")
	//timestampDemo2(2019-04-24  )

	fmt.Println("formatDemo.------->")
	formatDemo()
}
