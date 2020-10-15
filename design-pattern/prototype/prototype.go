package prototype

/*
如果对象的创建成本比较大，而同一个类的不同对象之间差别不大（大部分字段都相同），
在这种情况下，我们可以利用对已有对象（原型）进行复制（或者叫拷贝）的方式来创建新对象，
以达到节省创建时间的目的。这种基于原型来创建对象的方式就叫作原型设计模式（Prototype Design Pattern），
简称原型模式。
*/

import (
	"encoding/json"
	"time"
)

// Keyword 搜索关键字
type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

// Clone 这里使用序列化与反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
	var newKeyword Keyword
	b, _ := json.Marshal(k)
	json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

// Keywords 关键字 map
type Keywords map[string]*Keyword

// Clone 复制一个新的 keywords
// updatedWords: 需要更新的关键词列表，由于从数据库中获取数据常常是数组的方式
func (words Keywords) Clone(updatedWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		// 这里是浅拷贝，直接拷贝了地址
		newKeywords[k] = v
	}

	// 替换掉需要更新的字段，这里用的是深拷贝
	for _, word := range updatedWords {
		newKeywords[word.word] = word.Clone()
	}

	return newKeywords
}