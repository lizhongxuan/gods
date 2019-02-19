package main

import "fmt"

func quickSort(arr []int,l int,r int)  {
	if l>=r {
		return
	}
	i:=l
	j:=r
	key:=arr[i]  //选择第一个数为比较的key,
	for i<j  {
		for i<j && arr[j] > key{ //从右向左找到第一个小于key的值
			j--
		}
		if i<j {
			arr[i] = arr[j]
			i++
		}

		for i<j && arr[i]<key  { //从左向右找到第一个大于key的值
			i++
		}
		if i<j {
			arr[j] = arr[i]
			j--
		}

	}

	arr[i] = key
	quickSort(arr,l,i-1)
	quickSort(arr,i+1,r)
}

func main() {
	arr := []int{1, 7, 3, 5, 3, 6}
	quickSort(arr,0,len(arr)-1)
	fmt.Println(arr)
}
