package main

type node struct {
	pre  *node
	next *node
	t *task
	cycle_num int
}

type cycleList struct {
	taskList [3600]*node
	curIndex int
}

func NewCycleList()*cycleList  {
	var taskList [3600]*node
	return &cycleList{
		taskList:taskList,
		curIndex:0,
	}

}

