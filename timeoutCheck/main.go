package main

import (
	"time"
	"fmt"
)

func main()  {
	cl:=NewTimerCycle()
	go cl.Start()
	cl.addTask(5,"i")
	cl.addTask(5,"i2")
	cl.addTask(5,"i3")
	cl.addTask(6,"love")
	cl.addTask(7,"u")
	<-time.After(10 * time.Second)
	cl.Stop()

	<-time.After(3 * time.Second)

	fmt.Println("end")

}