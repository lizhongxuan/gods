package main

import (
	"DesignPattern/StatePattern/state"
	"time"
)

/*
状态模式把对象每一个状态的行为封装在对象内部。避免大量状态逻辑杂糅。
*/

func stateMechine(state state.State, ch chan int) {
	for {
		select {
		case i := <-ch:
			if i == 1 {
				state = state.NextState()
			}else if i == 0 {
				return
			}
		default:
			state.Update()
			time.Sleep(time.Second )

		}
	}

}

func main() {
	st := new(state.GameStartState)
	ch := make(chan int)
	go stateMechine(st, ch)
	time.Sleep(time.Second * 2)
	ch <- 1
	time.Sleep(time.Second * 2)
	ch <- 1
	time.Sleep(time.Second * 2)
	ch <- 0
	time.Sleep(time.Second * 2) }
