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
	for i:=0;i<3600 ;i++  {
		taskList[i]=&node{}
	}

	return &cycleList{
		taskList:taskList,
		curIndex:0,
	}

}

