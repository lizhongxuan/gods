package main

import (
	"fmt"
	"time"
	"sync"
)

type timerCycle struct {
	tasks   *cycleList
	taskMap map[*task]int
	isClose bool
	sync.Mutex
}

func NewTimerCycle() *timerCycle {
	return &timerCycle{
		tasks:   NewCycleList(),
		taskMap: make(map[*task]int),
		isClose: false,
	}
}

func (tc *timerCycle) addTask(second int, name string) {
	tc.Lock()
	defer tc.Unlock()
	cycle_num:= second/3600
	index := second%3600
	fmt.Println("addTask:",index)
	node_t := tc.tasks.taskList[index]

	newNode := &node{
		cycle_num:cycle_num,
		t:&task{
			name:name,
		},
	}

	if node_t == nil {
		tc.tasks.taskList[index] = newNode
		return
	}else if node_t.next == nil {
		node_t.next = newNode
		newNode.pre = node_t
		return
	}

	newNode.next = node_t.next
	node_t.next.pre = newNode
	node_t.next = newNode
	newNode.pre = node_t

}

func (tc *timerCycle) Start() {
	for {
		fmt.Println("index:",tc.tasks.curIndex)
		if tc.isClose {
			fmt.Println("close timerCycle")
			break
		}
		tc.checkTasks()

		if tc.tasks.curIndex == 3600 {
			tc.tasks.curIndex = -1
		}
		tc.tasks.curIndex++
		time.Sleep(1000 * time.Millisecond)
	}
}

func (tc *timerCycle)checkTasks()  {
	tc.Lock()
	defer tc.Unlock()

	indexNode := tc.tasks.taskList[tc.tasks.curIndex]
	tmpNode := indexNode

	for tmpNode != nil {
		if tmpNode.cycle_num == 0 {
			go tmpNode.t.print()
			if tmpNode.next ==nil {
				tmpNode = nil
				continue
			}

			tmpNode = tmpNode.next
			continue
		}
		tmpNode.cycle_num--
		tmpNode = tmpNode.next
	}
}

func (tc *timerCycle) Stop() {
	tc.isClose = true
}
