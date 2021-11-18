/*
练习：
	使用 random 生成随机数：
		伪随机数，根据一定的算法公式算出来的。
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 生成一个随机数：如果种子数不变，每次运行得到的随机数都是同一个数
	num1 := rand.Int()
	fmt.Println(num1)

	// 生成随机数，限定个数和范围：如果种子数不变，每次运行得到的随机数都是一样的
	for i := 0; i < 10; i++ {
		num := rand.Intn(10)
		fmt.Println(num)
	}

	// 通过修改种子数，即 rand.Seed() 里面的参数的数值，可以让每次运行得到的随机数不同
	rand.Seed(1000)
	num2 := rand.Intn(10)
	fmt.Println("-->", num2)

	// 怎样让种子数不停的改变，以便每次执行生成的随机数都不相同呢？即如何获取真正的随机数？
	/*
		思路：
			1、一直在变化的东西是什么？时间！
			2、如何获取时间？使用 time.Now()
			3、获取到的时间是什么格式？怎样转化为整数给 rand.Seed() 使用？Seed() 的参数要求是 int64 的整数
	*/
	t1 := time.Now()
	fmt.Println(t1)
	fmt.Printf("%T\n", t1) // 看 ti 是什么类型：是一个 time 对象
	// 时间戳：指定的某个时间距离 1970年1月1日0点0分0秒，之间的时间差值：秒，纳秒
	timeStamp1 := t1.Unix() // 秒
	fmt.Println(timeStamp1)

	timeStamp2 := t1.UnixNano() // 纳秒
	fmt.Println(timeStamp2)

	// 生成真正随机的过程：
	// step1：设置种子数，可以设置成时间戳
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		// step2：调用生成随机数的函数
		fmt.Println("-->", rand.Intn(100))
	}
}
