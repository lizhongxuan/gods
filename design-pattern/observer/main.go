package main

import "DesignPattern/ObserverPattern/observer"

//观察者模式使得一种类型的实例可以发送事件给其他类型，前提是接收事件的实例要根发送者订阅这个事件。


func main()  {
	eventCenter := observer.NewEventCenter()
	r_1 := observer.EventReciver{}
	r_2 := observer.EventReciver{}

	eventCenter.Register(&r_1)
	eventCenter.Register(&r_2)
	eventCenter.Notify(observer.Event{1})

	eventCenter.Degister(&r_1)
	eventCenter.Notify(observer.Event{2})
}
