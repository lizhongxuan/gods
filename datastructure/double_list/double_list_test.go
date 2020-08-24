package double_list

import (
	"testing"
	"fmt"
)

func TestDoubleList(t *testing.T) {
	l := NewDoubleList()
	l.InsertFront("1")
	l.InsertFront("asdf")
	l.InsertFront("22ff")
	l.InsertFront("333f")
	l.InsertFront("uuuuu")
	l.InsertBack("bbbbb")

	fmt.Println(l.seeFirst())
	fmt.Println(l.seeLast())


	l.PushFront()
	l.PushBack()

	fmt.Println(l.seeFirst())
	fmt.Println(l.seeLast())


	fmt.Println(l.len)
	t.Parallel()
}