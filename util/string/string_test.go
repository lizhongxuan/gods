package string

import (
	"testing"
	"fmt"
)

func Test_ToBool(t *testing.T)   {
	b,err:=ToBool("0")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)


	b,err=ToBool("1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)


	b,err=ToBool("2")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Test_ToInt32(t *testing.T)   {
	i,err := ToInt32("9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}

func Test_JoinSep(t *testing.T)   {
	str := JoinSep([]string{"123","456","789"},"&")
	fmt.Println(str)
}

func Test_FormatByte(t *testing.T)   {
	str := FormatByte(ToByte("李仲玄"))
	fmt.Println(str)
}
