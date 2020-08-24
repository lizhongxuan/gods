package main

import (
	"time"
	"fmt"
)

func main()  {
	cl:=NewTimerCycle()
	go cl.Start()
	cl.addTask(3615,"x1")
	cl.addTask(3,"i")
	cl.addTask(3,"i2")
	cl.addTask(3,"i3")
	cl.addTask(3,"i4")
	cl.addTask(3,"i5")
	cl.addTask(15,"x2")
	cl.addTask(15,"x3")
	cl.addTask(15,"x4")
	cl.addTask(6,"love")
	cl.addTask(7,"u")
	<-time.After(1000 * time.Second)
	cl.Stop()

	<-time.After(3 * time.Second)

	fmt.Println("end")

}