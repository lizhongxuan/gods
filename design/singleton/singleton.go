/*
@Time : 2020-07-11 10:28 
@Author : zhongxuanli
@File : singleton
@Software: GoLand
*/

package singleton

import "sync"

func main() {

}

type student struct {
	Name string
	Class string
}

var stu *student
var l sync.Mutex

//双重检查
func getStudent()*student  {
	if stu == nil {
		l.Lock()
		defer l.Unlock()
		if stu == nil {
			stu = &student{}
		}
	}
	return stu
}


var once sync.Once
//once.do
func getStudent2()*student  {
	once.Do(func() {
		stu = &student{}
	})
	return stu
}