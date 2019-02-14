package main

import "fmt"
func main() {
	var tree BPlusTree
	(&tree).Initialize()
	var i int
	i = 1
	for i < 100 {
		_, result := tree.Insert(i, i*10)
		fmt.Print(i)
		if result == false {
			print("数据已存在")
		}
		i++
	}
	tree.Remove(7)
	tree.Remove(6)
	tree.Remove(5)
	resultDate, success := tree.FindData(5)
	if success == true {
		fmt.Print(resultDate)
		fmt.Printf("\n")
	}

	//遍历结点元素
	i = 0
	for i < tree.root.Children[1].KeyNum {
		fmt.Println(tree.root.Children[1].leafNode.datas[i])
		i++
	}
}
