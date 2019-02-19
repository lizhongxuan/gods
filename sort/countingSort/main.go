package main

import "fmt"

func countingSort(arr []int)  {
	len := len(arr)
	if len<2 {
		return
	}

	var maxValue int = arr[0]
	//找到最大值
	for i:=1;i<len ;i++  {
		if maxValue < arr[i] {
			maxValue = arr[i]
		}
	}

	var bucket = make([]int,maxValue + 1)
	var sortedIndex int
	bucketLen := maxValue + 1

	//存入数组
	for i:=0;i<len ;i++  {
		bucket[arr[i]]++
	}

	//取出数组
	for i:=0;i<bucketLen ;i++ {
		for bucket[i] >0   {
			arr[sortedIndex] = i
			sortedIndex++
			bucket[i]--
		}
	}

}

func main() {
	arr := []int{1, 7, 3, 5, 3, 6}
	countingSort(arr)
	fmt.Println(arr)
}
