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

func (tc *timerCycle) addTask(afterTime int, name string) {
	tc.Lock()
	defer tc.Unlock()
	cycle_num:= afterTime / 3600
	index := afterTime % 3600
	fmt.Println("addTask:",index)
	node_t := tc.tasks.taskList[index]

	newNode := &node{
		cycle_num:cycle_num,
		t:&task{
			name:name,
		},
	}

	if node_t.next == nil {
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
		if tc.isClose {
			fmt.Println("close timerCycle")
			break
		}
		tc.checkTasks()

		if tc.tasks.curIndex == 3599 {
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
	tmpNode := indexNode.next
	for tmpNode != nil {
		//还没能触发,减一等下一轮再判断
		if tmpNode.cycle_num != 0 {
			tmpNode.cycle_num--
			tmpNode = tmpNode.next
			continue
		}
		//触发
		tmpNode.t.print()
		tmpNode.pre.next = tmpNode.next

		if tmpNode.next != nil {
			tmpNode.next.pre = tmpNode.pre
			tmpNode.pre.next = tmpNode.next
		}else {
			tmpNode.pre.next = nil
		}

		tmpNode = tmpNode.next
	}
}

func (tc *timerCycle) Stop() {
	tc.isClose = true
}
