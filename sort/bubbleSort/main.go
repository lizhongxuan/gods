package main

import "fmt"
//冒泡排序
//稳定算法,复杂度O(n^2)
func bubbleSort(arr []int)  {
	len:= len(arr)
	for i:=0;i<len-1 ;i++  {
		for j:=0;j<len-1 ;j++  {
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
}

func main() {
	arr := []int{1,7,3,5,3,6}
	bubbleSort(arr)
	fmt.Println(arr)
}