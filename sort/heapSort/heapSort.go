package main

import "fmt"

func main() {
	arr := []int{1, 7, 4, 5, 3, 6,2,9,12}
	minHeapSort(arr)
	fmt.Println(arr)
}

func minHeapSort(arr []int) {
	len := len(arr)
	//建立最小堆
	for i := (len - 1) / 2; i >= 0; i-- {
		minHeapFixdown(arr, i, len)
		fmt.Println("创建最小堆中:", arr)
	}

	fmt.Println("创建最小堆完成:", arr)

	//堆排序
	for i := len - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		minHeapFixdown(arr, 0, i)
		fmt.Println(arr)
		printPoint(i)


	}

	fmt.Println("堆排序完成:", arr)
}

func printPoint(num int)  {
	fmt.Print("--")
	for i:=0;i<num ;i++  {
		fmt.Print("--")
	}
	fmt.Println("↑")
}


//最小堆堆调整
func minHeapFixdown(arr []int, node, len int) {
	child := 2*node + 1 //取左子节点
	for child < len {

		//两个子节点取最小的
		if child+1 < len && arr[child+1] < arr[child] {
			child++
		}

		//节点小于子节点则退出
		//若大于,则与子节点交换
		if arr[node] <= arr[child] {
			break
		}
		arr[node], arr[child] = arr[child], arr[node]

		//往下层推进,判断子节点是否符合同样规律
		node = child
		child = 2*node + 1
	}
}
