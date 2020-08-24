package main

import (
	"DesignPattern/MementoPattern/memento"
	"fmt"
)

/*
  备忘录模式保存一个对象的状态，在需要的时候将其恢复。
  该模式在不破坏封装的前提下，捕获一个对象的内部状态，
  并在该对象之外保存这个状态，这样可以在以后将对象恢复到原先保存的状态。
  很多时候我们总是需要记录一个对象的内部状态，这样做的目的就是为了允许用户取消不确定或者错误的操作，
  能够恢复到他原先的状态，使得他有”后悔药”可吃。
*/
/*
应用实例
    后悔药
    打游戏时的存档
    Windows 里的 ctri + z
    IE 中的后退
    数据库的事务管理

优点
    给用户提供了一种可以恢复状态的机制，可以使用户能够比较方便地回到某个历史的状态
    实现了信息的封装，使得用户不需要关心状态的保存细节

缺点
    消耗资源。如果类的成员变量过多，势必会占用比较大的资源，而且每一次保存都会消耗一定的内存。

*/
func main() {
	s:=memento.NewCaretakerRoleMemory()
	man := memento.NewRole(100)

	//存档
	s.Save(man.Save())

	//战斗
	man.Fight()
	fmt.Print("战斗 当前状态:")
	fmt.Println(*man)

	//再存档
	s.Save(man.Save())

	//再战斗
	man.Fight()
	fmt.Print("再战斗 当前状态:")
	fmt.Println(*man)

	//回档
	man.Read(s.GetAndRemoveMemory())
	fmt.Print("回档 当前状态:")
	fmt.Println(*man)

	//再回档
	man.Read(s.GetAndRemoveMemory())
	fmt.Print("再回档 当前状态:")
	fmt.Println(*man)
}