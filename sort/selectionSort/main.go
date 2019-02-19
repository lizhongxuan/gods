package main

import "fmt"

//选择排序
//不稳定算法,时间复杂度O(n^2)
func selectionSort(arr []int) {
	var len = len(arr)
	var minIndex int
	for i := 0; i < len-1; i++ {
		minIndex = i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] { // 寻找最小的数
				minIndex = j // 将最小数的索引保存
			}
		}
		arr[i],arr[minIndex] = arr[minIndex],arr[i]
	}
}

func main() {
	arr := []int{1, 7, 3, 5, 3, 6}
	selectionSort(arr)
	fmt.Println(arr)
}
