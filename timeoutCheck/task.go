package main

import "fmt"

type task struct {
	name  string
}

func (t *task)print()  {
	fmt.Println("我是",t.name,",我被触发!!!")
}



