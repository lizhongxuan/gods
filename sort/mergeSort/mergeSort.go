package main

import "fmt"

func main() {
	arr := []int{1, 7, 3, 5, 3, 6}
	fmt.Println(mergeSort(arr))
}
func mergeSort(arr []int) []int {
	var len = len(arr)
	if len<2 {
		return arr
	}
	var middle  =  len/2
	left := arr[:middle]
	right := arr[middle:]
	return merge(mergeSort(left),mergeSort(right))
}

func merge(leftArr, rightArr []int) []int {
	var result []int
	var left,right int
	for len(leftArr)>left && len(rightArr) >right {
		if leftArr[left] <= rightArr[right] {
			result = append(result,leftArr[left])
			left++
		}else {
			result = append(result,rightArr[right])
			right++
		}
	}
	for len(leftArr)>left{
		result = append(result,leftArr[left])
		left++
	}
	for len(rightArr) >right{
		result = append(result,rightArr[right])
		right++
	}
	fmt.Println(result)
	return result
}
