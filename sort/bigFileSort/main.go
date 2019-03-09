package main

import (
	"fmt"
	"math"
)

const MaxNumber = 9999

// 声明一个节点的结构体，包含节点数值大小和是否需要参与比较
type node struct {
	// 数值大小
	value int
	// 叶子节点状态
	available bool
	// 叶子中的排序，方便失效
	group_id int
}

func main() {
	//代表大文件
	bigArr := []int{1, 54, 3, 7, 4, 5, 7, 8, 12, 9, 7, 6, 34, 123, 5, 7, 98, 235, 34, 76, 7, 2, 5, 7, 25, 36}
	splitNum := 5

	//切分为多个小文件
	smallArrs := splitArr(bigArr, splitNum)
	fmt.Println("smallArrs:", smallArrs)

	//小文件排序
	for _, small := range smallArrs {
		quickSort(small, 0, len(small)-1)
	}
	fmt.Println("smallArrs 排序后:", smallArrs)

	//小文件的数量
	sls := len(smallArrs)

	//每个小文件输出的数量
	var arrNumbers = make([]int, sls)

	//获取每个文件第一个数,也是每个文件的最小值,建立败者树
	var minArr []int
	for group_id := 0; group_id < sls; group_id++ {
		minArr = append(minArr, smallArrs[group_id][0])
	}

	//合并最终存放文件
	var endArr []int

	//创建胜者树
	tree, level := createTree(minArr)

	//循环一次,取一个最小值
	for tree[0].value != MaxNumber {
		endArr = append(endArr, tree[0].value)

		fmt.Println("tree[0].group_id:", tree[0].group_id)
		group_id := tree[0].group_id
		sl := len(smallArrs[group_id])

		value := MaxNumber

		if arrNumbers[group_id]+1 < sl {
			arrNumbers[group_id]++ //组输出的数量加一
			value = smallArrs[group_id][arrNumbers[group_id]]
		}

		//胜者树排序
		winTreeSort(tree, level,value)
	}
	fmt.Println("endArr:",endArr)
}

func winTreeSort(tree []node, level int,value int) {
	var leaf = pow(2, level)
	winNode := tree[0].group_id + leaf - 1

	tree[winNode].value = value

	for i := 0; i < level; i++ {
		leftNode := winNode
		if winNode%2 == 0 {
			leftNode = winNode - 1
		}

		// 比较兄弟节点间大小，并将胜出的节点向上传递
		compareAndUp(&tree, leftNode)
		winNode = (leftNode - 1) / 2
	}
}

func createTree(origin []int) ([]node, int) {
	// 树的层数
	var level int
	for pow(2, level) < len(origin) {
		level++
	}
	// 叶子节点数
	var leaf = pow(2, level)
	var tree = make([]node, leaf*2-1)

	// 先填充叶子节点的数据
	for i := 0; i < len(origin); i++ {
		tree[leaf+i-1] = node{origin[i], true, i}
	}

	// 每层都比较叶子兄弟大小，选出较大值作为父节点
	for i := 0; i < level; i++ {
		// 当前层节点数
		nodeCount := pow(2, level-i)
		// 每组兄弟间比较
		for j := 0; j < nodeCount/2; j++ {
			compareAndUp(&tree, nodeCount-1+j*2)
		}
	}

	// 这个时候树顶端的就是最小的元素了
	return tree, level
}

func compareAndUp(tree *[]node, leftNode int) {
	rightNode := leftNode + 1
	// 除非左节点无效或者右节点有效并且比左节点大，否则就无脑选左节点
	if !(*tree)[leftNode].available || ((*tree)[rightNode].available && (*tree)[leftNode].value > (*tree)[rightNode].value) {
		(*tree)[(leftNode-1)/2] = (*tree)[rightNode]
	} else {
		(*tree)[(leftNode-1)/2] = (*tree)[leftNode]
	}
}

func pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}


//切分文件
func splitArr(arr []int, splitNum int) [][]int {
	l := len(arr)
	var reqArr [][]int
	if l == 0 || splitNum < 1 {
		return reqArr
	} else if splitNum == 1 {
		reqArr = append(reqArr, arr)
		return reqArr
	}

	for i := 0; i < l; i += splitNum {
		if l-i <= splitNum {
			reqArr = append(reqArr, arr[i:])
			return reqArr
		}
		reqArr = append(reqArr, arr[i:i+splitNum])
	}
	return reqArr
}

//排序
func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	i := l
	j := r
	key := arr[i] //选择第一个数为比较的key,
	for i < j {
		for i < j && arr[j] > key { //从右向左找到第一个小于key的值
			j--
		}
		if i < j {
			arr[i] = arr[j]
			i++
		}

		for i < j && arr[i] < key { //从左向右找到第一个大于key的值
			i++
		}
		if i < j {
			arr[j] = arr[i]
			j--
		}
	}
	arr[i] = key
	quickSort(arr, l, i-1)
	quickSort(arr, i+1, r)
}
