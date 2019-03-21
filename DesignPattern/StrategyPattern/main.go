package main

import (
	"DesignPattern/StrategyPattern/strategy"
	"fmt"
)

/*
策略模式可以让更换对象的内脏，而装饰者模式可以更换对象的外表。
*/


func main() {
	operation := strategy.Operation{strategy.Addition1{}}
	res  := operation.Operate(3,3)
	fmt.Println(res)

	operation2 := strategy.Operation{strategy.Addition2{}}
	res2  := operation2.Operate(3,3)
	fmt.Println(res2)
}