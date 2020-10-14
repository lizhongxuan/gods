package main

type Event struct {
	Data int64
}

type Observer interface {
	OnNotify(Event)
}

type Notifier interface {
	Register(Observer)
	Degister(Observer)
	Notifier(Event)
}
