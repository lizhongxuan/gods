package main

import "fmt"

//插入排序
//稳定算法,时间复杂度O(n^2)
func insetionSort(arr []int)  {
	len:=len(arr)
	var preIndex,current int
	for i:=1;i<len ;i++  {
		preIndex = i-1
		current = arr[i]
		for preIndex >= 0 && arr[preIndex] > current  {
			arr[preIndex+1] = arr[preIndex]
			preIndex--
		}
		arr[preIndex+1] = current
	}
}
func main() {
	arr := []int{1, 7, 3, 5, 3, 6}
	insetionSort(arr)
	fmt.Println(arr)
}