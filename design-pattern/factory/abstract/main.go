package main

import "fmt"
/*
抽象工厂

  抽象工厂是针对一个产品族而言的.
  产品族就好像套餐.一个套餐包含了好几种食品，而每一种食品都是一种类型的食物。举个例子.

  一个套餐定义为食品和饮料.而有一家餐馆的食品包括肉和汉堡，饮料包括CoCo和茶.
  这家餐馆不想单独出售某种食品和饮料，只卖套餐.

  于是老板定义:
    套餐A:肉和CoCo
    套餐B:汉堡和茶

  定好了之后,把套餐A外包给工厂A负责生产，套餐B外包给工厂B负责生产。
  两个工厂根据这家店的需求，实习那了生产食品和生产饮料的方法。
  A工厂就负责A套餐，那么他就需要实现生产肉和CoCo的逻辑即可.
  而B工厂只需要实现生产汉堡和茶的逻辑即可.

  这样以来，来到店里的客人，只需要订购套餐，服务员通知工厂生产并送达即可。
  假设一个客人要套餐A，服务员通知工厂A，先生产一个肉，再来一杯CoCo，
  服务员负责把这些产品递给客人食用即可。

*/



// 创建的接口
type Factory interface {
	CreateFood() Food
	CreateDrink() Drink
}
func NewFactory(name string)Factory  {
	switch name {
	case "a":
		return FactoryA{}
	case "b":
		return FactoryB{}
	}
	return nil
}

// 抽象层实现
type FactoryA struct{}
func (af FactoryA) CreateFood() Food {
	f := Meat{}
	return f
}
func (af FactoryA) CreateDrink() Drink {
	d := CoCo{}
	return d
}


type FactoryB struct{}
func (bf FactoryB) CreateFood() Food {
	f := Hamberger{}
	return f
}
func (bf FactoryB) CreateDrink() Drink {
	d := Tea{}
	return d
}


// 食物使用的接口
type Food interface {
	Eat()
}
type Meat struct{}
func (m Meat) Eat()      { fmt.Println("Eat meat.") }

type Hamberger struct{}
func (h Hamberger) Eat() { fmt.Println("Eat Hamberger.") }


type Drink interface{
	Drink()
}

type CoCo struct{}
func (cc CoCo) Drink() { fmt.Println("Drink CoCo") }

type Tea struct{}
func (t Tea) Drink() { fmt.Println("Drink Tea") }

func main() { // Abstract Factory
	fa := NewFactory("a")
	fa.CreateFood().Eat()
	fa.CreateDrink().Drink()


	fb := NewFactory("b")
	fb.CreateFood().Eat()
	fb.CreateDrink().Drink()
}
