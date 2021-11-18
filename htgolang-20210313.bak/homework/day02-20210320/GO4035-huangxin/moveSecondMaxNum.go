package main

/*
	create-date:2021-03-21
	describe: 移动第二大值到倒数第二位置
	author:GO4035-huangxin
*/

import "fmt"

func main() {
	/*
				找到切片中最大的元素索引和值(打印)
		    	[]int{9, 8, 19, 10, 2, 8}
			第一步：找到 最大值 19 并在切片中进行删除
			第二步：找打 最大值 10 就是第二大的值

	*/

	// 源切片内容
	nums := []int{9, 8, 19, 10, 2, 8}
	fmt.Println("源nums:", nums)

	// 通过方法找到第一大值的 index
	max_index, max_num := getMaxNum(nums)

	// 定义接收 第二大值与其索引
	secondMaxIndex := 0
	secondMaxNum := 0

	// 找到 nums 第二大值与其索引
	for k, v := range nums {
		if k == max_index && v == max_num {
			continue
		} else {
			if secondMaxNum < v {
				secondMaxNum = v
				secondMaxIndex = k
			}
		}
	}

	// 找到打印第二大值与其索引
	//fmt.Printf("nums 中第二大值为 %d，其索引为 %d\n", secondMaxNum, secondMaxIndex)

	// 在nums中先进行删除第二大值
	copy(nums[secondMaxIndex:], nums[secondMaxIndex+1:])
	nums = nums[:len(nums)-1]
	// fmt.Println(nums)

	// 将 第二大值追加到 nums中
	nums = append(nums, secondMaxNum)
	// fmt.Println(nums)
	// 将nums 最后一个元素 与 倒数第二个元素 进行交换
	nums[len(nums)-1], nums[len(nums)-2] = nums[len(nums)-2], nums[len(nums)-1]
	fmt.Println("移动后nums:", nums)

}

// 定义返回最大值的 index 方法
func getMaxNum(nums []int) (int, int) {

	max_num := 0
	max_index := 0

	for i, v := range nums {
		if max_num < v {
			max_num = v
			max_index = i
		}
	}

	return max_index, max_num
}
