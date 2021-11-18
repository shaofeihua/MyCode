// 使用 for 循环实现猜数字游戏
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	endNum := 100
	i := rand.Intn(endNum)
	//fmt.Println(i)

	fmt.Println("猜数字游戏开始！请您输入 0-100 之间整数，您总共有 5 次猜测的机会：")
	reader := bufio.NewReader(os.Stdin)
	for a := 1; a <= 5; a++ {
		//fmt.Scan(&num)
		data, _, _ := reader.ReadLine()
		//fmt.Printf("%T\n", data)
		num, err := strconv.Atoi(string(data)) //string to int,并作输入格式判断
		if err != nil {
			fmt.Println("格式不对，请输入数字")
		} else if a == 5 && num != i {
			fmt.Println("对不起，5 次都没猜对。游戏结束~")
			fmt.Printf("正确的数字是：%d\n", i)
			break
		} else if num < i {
			fmt.Println("您输入的数字小了。请再次输入：")
		} else if num > i {
			fmt.Println("您输入的数字大了。请再次输入：")
		} else if num == i {
			fmt.Println("您猜对了！游戏结束~")
			break
		}
	}
}
