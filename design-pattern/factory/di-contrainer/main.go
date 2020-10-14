package main

import "fmt"

/*
 DI - Dependency Injection，意为依赖注射（依赖注入）。
指组件（COMPONENT：这里可以看作提供某功能的class或class群）之间依赖系统（容器）的注入（injection）而产生关联关系。
也被称为控制反转（Ioc：Inversion of Control）模式。
通俗一点讲，该模式将实现组件间关系从程序内部提到外部容器来管理。
*/

type FoodFactory struct{}
type Food interface{
	Eat()
}
type Hamberger struct{

}


func (ff FoodFactory) CreateFood(name string) Food {
	var s Food
	switch name {
	case "Hamberger":
		s = new(Hamberger)
	}
	return s
}

func (h Hamberger)Eat () {
	fmt.Println("Eat Hamberger.")
}



type A struct {
}

func (a A)Have(h Food)  {
	h.Eat()
}

func main() {
	f := FoodFactory{}
	h :=f.CreateFood("Hamberger")

	a := A{}
	a.Have(h)
}