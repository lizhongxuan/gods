package main

//go vet govet.go

import (
	"sync"
)

//omitempty表示在打印时若该项为empty则不打印，应将其放在双引号内
type Parameters struct {
	Unit        int `json:"test_unit"`
	MaxInstance int `json:"max_instance",omitempty` // struct field tag `json:"max_instance",omitempty` not compatible with reflect.StructTag.Get: key:"value" pairs not separated by spaces
	MinInstance int `json:"min_instance",omitempty` // struct field tag `json:"min_instance",omitempty` not compatible with reflect.StructTag.Get: key:"value" pairs not separated by spaces
}

//不能值传递锁，否则可能导致死锁
func createTest(message chan []byte, lock sync.Mutex) { //createTest passes lock by value: sync.Mutex

}

//建议将tag改为pair形式
type LoggerConfig struct {
	Level string "level" // struct field tag `level` not compatible with reflect.StructTag.Get: bad syntax for struct tag pair
	File  string "file"  // struct field tag `file` not compatible with reflect.StructTag.Get: bad syntax for struct tag pair
}
