package main

import "fmt"

/*
	工厂方法
	将复杂的创建逻辑拆分到多个工厂类中，让每个工厂类都不至于过于复杂
	例子: 将对象的创建跟使用拆分
*/

// 创建的接口
type Factory interface {
	Create() Food
}
func NewFactory(name string)Factory  {
	switch name {
	case "meat":
		return MeatFactory{}
	case "hamberger":
		return HambergerFactory{}
	}
	return nil
}

type MeatFactory struct{}
func (mf MeatFactory) Create() Food {
	m := Meat{}
	return m
}

type HambergerFactory struct{}
func (hf HambergerFactory) Create() Food {
	h := Hamberger{}
	return h
}


// 食物的接口
type Food interface {
	Eat()
}
type Meat struct{}
type Hamberger struct{}

func (m Meat) Eat() {
	fmt.Println("Eat meat.")
}
func (h Hamberger) Eat() {
	fmt.Println("Eat Hamberger.")
}

func main() {
	// Factory Method
	mf := NewFactory("meat")
	mf.Create().Eat()
	hf := NewFactory("hamberger")
	hf.Create().Eat()
}
