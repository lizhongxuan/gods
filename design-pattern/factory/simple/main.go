package main

import "fmt"

/***
当创建逻辑比较复杂，是一个“大工程”的时候，我们就考虑使用工厂模式，封装对象的创建过程，将对象的创建和使用相分离。

 *  Simple Factory
简单工厂的实现思想，即创建一个工厂，将产品的实现逻辑集中在这个工厂中。
由于 Go 本身是没有构造函数的，一般而言我们采用 NewName 的方式创建对象/接口，当它返回的是接口的时候，其实就是简单工厂模式
 */
type FoodFactory struct{}
type Food interface{
	Eat()
}
type Meat struct{}
type Hamberger struct{}


func (ff FoodFactory) CreateFood(name string) Food {
	var s Food
	switch name {
	case "Meat":
		s = new(Meat)
	case "Hamberger":
		s = new(Hamberger)
	}
	return s
}

func (m Meat) Eat() {
	fmt.Println("Eat meat.")
}
func (h Hamberger) Eat() {
	fmt.Println("Eat Hamberger.")
}

func main() {
	// Simple Factory
	f := FoodFactory{}
	f.CreateFood("Meat").Eat()
	f.CreateFood("Hamberger").Eat()
}
