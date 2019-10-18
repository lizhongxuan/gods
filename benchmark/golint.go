package main

//golint golint.go
/*
变量名规范
变量的声明，像var str string = "test"，会有警告，应该var str = "test"
大小写问题，大写导出包的要有注释
x += 1 应该 x++
....
*/

import "fmt"
import "errors"

//先定义几个类型 //这行注释应该是关于Color的注释  comment on exported type Color should be of the form "Color ..." (with optional leading article)
type Color string
type LampStatus bool // 导出的参数(首字母大写)应该有注释 exported type LampStatus should have comment or be unexported
type Brand string    //exported type Brand should have comment or be unexported

// 定义颜色常量
const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

// 定义品牌常量
const (
	OppleBulb Brand = "OPPLE"
	Osram           = "OSRAM"
)

//  Lamp Builder define
type Builder interface {
	Color(Color) LampBuilder
	Brand(Brand) LampBuilder
	Build() LampOperation
}

type LampBuilder struct {
	Lamp // 配置结构
}

func (lb LampBuilder) Color(c Color) LampBuilder {
	lb.color = c
	return lb
}

func (lb LampBuilder) Brand(b Brand) LampBuilder {
	lb.brand = b
	return lb
}

func (lb LampBuilder) Build() LampOperation {
	// 新的产品产生过程
	lamp := Lamp{color: lb.color, brand: lb.brand, status: false}
	return lamp
}

func NewBuilder() Builder { return LampBuilder{} }

type LampOperation interface {
	Open() error
	Close() error
	ProductionIllustrative()
}

// 灯的定义
type Lamp struct {
	color  Color
	brand  Brand
	status LampStatus
}

func (l Lamp) Open() error {
	if l.status {
		return errors.New("Lamp is already opened")
	}
	fmt.Println("Open lamp.")
	l.status = true
	return nil
}

func (l Lamp) Close() error {
	if !l.status {
		return errors.New("Lamp is closed")
	}
	fmt.Println("Close lamp.")
	l.status = true
	return nil
}

func (l Lamp) ProductionIllustrative() {
	fmt.Println("I'm a lamp.")
	fmt.Println("Color:" + l.color)
	fmt.Println("Brand:" + l.brand)
}
