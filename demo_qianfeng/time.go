package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 获取当前时间
	t1 := time.Now()
	fmt.Printf("t1 的类型是：%T，t1 的值是：%v\n", t1, t1)

	// 获取指定的时间
	t2 := time.Date(2020, 8, 8, 12, 18, 36, 0, time.Local)
	fmt.Printf("t2 的类型是：%T，t2 的值是：%v\n", t2, t2)

	// 按照给出的字符串格式，输出当前的时间
	// 2006-1-2 15:04:05 是 golang 诞生的时间。必须使用这个时间才能格式化输出
	s1 := t1.Format("2006-1-2 15:04:05")
	fmt.Printf("s1 的类型是：%T，s1 的值是：%v\n", s1, s1)
	s2 := t1.Format("2006/01/02")
	fmt.Printf("s2 的类型是：%T，s2 的值是：%v\n", s2, s2)

	// 将 string 转化为与模板格式相同的时间
	// 2006-1-2 15:04:05 是 golang 诞生的时间。必须使用这个时间才能格式化输出
	s3 := "1999年10月10日"                      // string
	t3, err := time.Parse("2006年01月02日", s3) // 不能使用 2006/01/02 或者 2006-01-02
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Printf("t3 的类型是：%T，t3 的值是：%v\n", t3, t3)

	fmt.Println(t1.String())

	// 根据当前时间获取指定的内容
	// 获取当前的年、月、日
	year, month, day := t1.Date()
	fmt.Println(year, month, day)
	// 获取当前的时、分、秒
	hour, min, sec := t1.Clock()
	fmt.Println(hour, min, sec)
	// 仅获取年。仅获取月、日、时、分、秒、微秒、纳秒的方式类同。
	year2 := t1.Year()
	fmt.Println(year2)
	// 今年已经过了多少天
	fmt.Println(t1.YearDay())
	// 获取星期
	fmt.Println(t1.Weekday())

	// 时间戳：制定的日期，距离1970年1月1日0时0分0秒的时间差值：秒，纳秒
	t4 := time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC)
	timeStamp1 := t4.Unix() // 计算时间 t4 距离1970年1月1日0时0分0秒的时间差值
	fmt.Println(timeStamp1)
	// 求当前时间的时间戳
	timeStamp2 := t1.Unix()
	fmt.Println(timeStamp2)
	// 求当前时间的时间戳：纳秒
	timeStamp3 := t1.UnixNano()
	fmt.Println(timeStamp3)

	// 时间间隔
	fmt.Println("当前时间是：", t1)
	fmt.Println("当前时间加上 1 分钟：", t1.Add(time.Minute))
	fmt.Println("当前时间加上 1 天：", t1.Add(time.Hour*24))
	fmt.Println("当前时间加上 1 年：", t1.AddDate(1, 0, 0)) // AddDate(year,month,day)
	// 计算两个时间的时间间隔(差值)
	d1 := t4.Sub(t1)
	fmt.Println(d1) // -443820h36m27.0520695s
	d2 := t1.Sub(t4)
	fmt.Println(d2) // 443820h36m27.0520695s

	// 睡眠
	time.Sleep(time.Second * 3) // 睡眠 3 秒钟
	fmt.Println("go on ...")

	// 睡眠[1-10]的随机秒数
	rand.Seed(time.Now().UnixNano()) // 获取种子
	randNum := rand.Intn(10) + 1     // 获得 1-10 的随机数。如果不加 1 ，则范围为 0-9
	fmt.Println("随机数为：", randNum)
	time.Sleep(time.Duration(randNum) * time.Second)
	fmt.Println("go on ...")
}
