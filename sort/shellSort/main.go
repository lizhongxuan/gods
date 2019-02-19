package main

import "fmt"

func shellSort(arr []int)  {
	len := len(arr)
	var temp int
	var gap  int = 1
	for gap<len/3  {
		gap = gap * 3 + 1
	}

	for gap>0  {

		for i:=gap;i<len ;i++  {
			temp = arr[i]
			var j int
			for j =i-gap; j>0 && arr[j]>temp; j -= gap {
				arr[j+gap] = arr[j]
			}
			arr[j+gap] = temp
		}
		gap = gap/3
	}
}

func main() {
	arr := []int{1, 7, 3, 5, 3, 6}
	shellSort(arr)
	fmt.Println(arr)
}
