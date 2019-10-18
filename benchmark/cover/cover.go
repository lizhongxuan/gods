package cover

import "fmt"

type HJXFactory interface {
	CreateFood() Food
	CreateDrink() Drink
}

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

type Food interface {
	Eat()
}

type Meat struct{}
type Hamberger struct{}

func (m Meat) Eat()      { fmt.Println("Eat meat.") }
func (h Hamberger) Eat() { fmt.Println("Eat Hamberger.") }

type Drink interface{ Drink() }
type CoCo struct{}

func (cc CoCo) Drink() { fmt.Println("Drink CoCo") }

type Tea struct{}

func (t Tea) Drink() { fmt.Println("Drink Tea") }

func cover_demo(score int) string {
	switch {
	case score < 60:
		return "D"
	case score <= 70:
		return "C"
	case score <= 80:
		return "B"
	case score <= 90:
		return "A"
	default:
		return "Undefined"
	}
}
