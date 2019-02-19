package main

import "fmt"

func main()  {
	arr := []int{1, 7, 3, 5, 3, 6}
	minHeapSort(arr)
	fmt.Println(arr)
}

func minHeapSort(arr []int) {
	len := len(arr)
	//建立最小堆
	for i := (len - 1) / 2; i >= 0; i-- {
		minHeapFixdown(arr, i, len)
	}

	//堆排序
	for i := len - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		minHeapFixdown(arr, 0, i)
	}
}

//最小堆堆调整
func minHeapFixdown(arr []int, i, len int) {
	j := 2*i+1  //取左子节点
	for j < len {

		//两个子节点取最小的
		if j+1 < len && arr[j+1] < arr[j] {
			j++
		}

		//节点小于子节点则退出
		//若大于,则与子节点交换
		if arr[i] <= arr[j] {
			break
		}
		arr[i],arr[j] = arr[j],arr[i]

		//往上层推进
		i=j
		j = 2*i +1
	}
}
