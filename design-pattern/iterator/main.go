/*
迭代器（iterator）是一种对象，它能够用来遍历标准模板库容器中的部分或全部元素，每个迭代器对象代表容器中的确定的地址。
迭代器修改了常规指针的接口，所谓迭代器是一种概念上的抽象：那些行为上像迭代器的东西都可以叫做迭代器。
然而迭代器有很多不同的能力，它可以把抽象容器和通用算法有机的统一起来。
*/


package main

import (
	"fmt"
)

func main() {
	aSet := iterSet{make(map[string]struct{})}
	aSet.Add("hello")
	aSet.Add("hello")
	aSet.Add("world")
	aSet.Add("world")

	iter := aSet.Iterator()

	for v := range iter.C {
		fmt.Printf("key: %s\n", v.(string))
	}
}

// Iterator 声明接口
type Iterator interface {
	Iterator(m iterSet) Iter
}

// Iter 迭代器的实现
type Iter struct {
	C chan interface{}
}

func newIter(i *iterSet) Iter {
	iter := Iter{make(chan interface{})}

	go func() {
		for k := range i.m {
			iter.C <- k
		}
		close(iter.C)
	}()

	return iter
}

// 我们自己的set
type iterSet struct {
	m map[string]struct{}
}

// Add 添加元素
func (i *iterSet) Add(k string) {
	i.m[k] = struct{}{}
}

// Iterator 返回一个迭代器
func (i *iterSet) Iterator() Iter {
	return newIter(i)
}
