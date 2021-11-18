/*
练习：求水仙花数
	100 - 999 以内的三位数
	这个数的每个位上的数字的立方的和，刚好等于这个数字本身，这个数就叫水仙花数
	比如：153
		1*1*1 + 5*5*5 + 3*3*3 = 153

	分析：如何从 153 中取到 1、5、3 这三个数？
	取百位数 1 ：153 / 100 = 1
	取十位数 5 ：153 / 10 % 10 = 5
	取个位数 3 ：153 % 10 = 3
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 100; i <= 999; i++ {
		a := i / 100
		b := i / 10 % 10
		c := i % 10
		/*
			注意：
				1、Pow 要求参数类型必须是 float64，所以必须先使用 float64() 将 a、b、c 转化
				2、由于 a、b、c 相加得到的结果应该也是 float64，所以 == 后面的 i 也必须使用 float64() 进行转化
				3、这里的打印要使用 Print 或者 Println 而不是 Printf，否则报错：cannot use i (type int) as type string in argument to fmt.Printf
		*/
		if math.Pow(float64(a), 3)+math.Pow(float64(b), 3)+math.Pow(float64(c), 3) == float64(i) {
			fmt.Print(i, "\t")
		}
	}

	fmt.Println("\n-------------------")
	/*
		另一种思路：
		百位数：1-9
		十位数：0-9
		个位数：0-9

		使用多层循环，将百位数的数字、十位数的数字、个位数的数字看成独立的数，轮询这些数字并从中选择满足条件的数字：
	*/
	for a := 1; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				n := a*100 + b*10 + c*1
				if a*a*a+b*b*b+c*c*c == n {
					fmt.Print(n, "\t")
				}
			}
		}
	}
}
