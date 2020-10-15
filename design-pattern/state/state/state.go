package state

import "fmt"

type State interface {
	NextState() State
	Update()
}
type GameStartState struct{}
func (this *GameStartState) NextState() State {
	fmt.Println("Start Next")
	return new(GameRunState)
}

func (this *GameStartState) Update() {
	fmt.Println("Game Start")
}


type GameRunState struct{}
func (this *GameRunState) NextState() State {
	fmt.Println("Run Next")
	return new(GameEndState)
}

func (this *GameRunState) Update() {
	fmt.Println("Game Run")
}



type GameEndState struct{}
func (this *GameEndState) NextState() State {
	return new(GameStartState)
}

func (this *GameEndState) Update() {
	fmt.Println("Game End")
}
