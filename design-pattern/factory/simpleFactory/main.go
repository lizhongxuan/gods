package main

import "fmt"

/***
 *  Simple Factory
简单工厂的实现思想，即创建一个工厂，将产品的实现逻辑集中在这个工厂中。
 */
type FoodFactory struct{}

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

type Food interface{ Eat() }
type Meat struct{}
type Hamberger struct{}

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
