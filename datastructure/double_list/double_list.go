package double_list

import (
	"sync/atomic"
	"fmt"
)

type node struct {
	pre *node
	next *node
	value string
}

type doubleList struct {
	firstNode *node
	lastNode *node
	len int64
}

func NewDoubleList()*doubleList  {
	return &doubleList{}
}

func (l *doubleList)seeFirst()string  {
	return l.firstNode.value
}

func (l *doubleList)seeLast()string  {
	return l.lastNode.value
}


func (l *doubleList)PushFront()string  {
	if l.firstNode ==nil {
		return ""
	}
	v := l.firstNode.value


	l.firstNode = l.firstNode.next
	l.firstNode.pre.next = nil
	l.firstNode.pre = nil


	atomic.AddInt64(&l.len,-1)
	return v
}

func (l *doubleList)PushBack()string  {
	if l.lastNode ==nil {
		return ""
	}
	v := l.lastNode.value

	l.lastNode = l.lastNode.pre
	l.lastNode.next.pre = nil
	l.lastNode.next = nil
	atomic.AddInt64(&l.len,-1)
	return v
}
func (l *doubleList)InsertFront(v string)  {
	fmt.Println("len:",l.len)
	if l.len==0 {
		n := &node{
			value:v,
		}
		atomic.AddInt64(&l.len,1)
		l.firstNode=n
		l.lastNode=n
		return
	}

	n := &node{
		next:l.firstNode,
		value:v,
	}
	l.firstNode.pre = n
	l.firstNode = n
	atomic.AddInt64(&l.len,1)
}


func (l *doubleList)InsertBack(v string)  {
	if l.len==0 {
		n := &node{
			value:v,
		}
		atomic.AddInt64(&l.len,1)
		l.firstNode=n
		l.lastNode=n
		return
	}

	n := &node{
		pre:l.lastNode,
		value:v,
	}
	l.lastNode.next = n
	l.lastNode = n
	atomic.AddInt64(&l.len,1)
}
