/*
@Time : 2020-07-11 10:28 
@Author : zhongxuanli
@File : singleton
@Software: GoLand
*/

package singleton
/***
 单例模式
单例存在哪些问题？
单例对 OOP 特性的支持不友好
单例会隐藏类之间的依赖关系
单例对代码的扩展性不友好
单例对代码的可测试性不友好
单例不支持有参数的构造函数
 */
import "sync"

type student struct {
	Name string
	Class string
}

var stu *student
var l sync.Mutex

/* 饿汉式 init阶段就初始化对象 */
func init() {
	stu = &student{}
}


/* 懒汉式 虽然支持延迟加载,但需要加锁 */
//双重检查
func getStudent()*student  {
	if stu == nil {
		l.Lock()
		defer l.Unlock()
		if stu == nil { //防止上一个竞争者,把stu置为nil
			stu = &student{}
		}
	}
	return stu
}



//once.do
var once sync.Once
func getStudent2()*student  {
	once.Do(func() {
		stu = &student{}
	})
	return stu
}