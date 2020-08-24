package main

import (
	"fmt"
)

/*
策略模式可以让更换对象的内脏，而装饰者模式可以更换对象的外表。
*/

func main() {
	//加法
	operation := Operation{Add{}}
	res  := operation.Operate(3,3)
	fmt.Println(res)

	//乘法
	operation2 := Operation{Multiplication{}}
	res2  := operation2.Operate(3,3)
	fmt.Println(res2)
}