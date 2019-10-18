package bench

import (
	"encoding/json"
	"sync"
	"sync/atomic"
	"testing"
)

//go test -cpu -bench=.
//如果有TestMain函数，函数中需要有t.Run(), 否则benchmark不能运行

type GetVarFunc func(key string) interface{}

func benchmark(b *testing.B, fn GetVarFunc) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fn("foo")
		}
	})
}

func BenchmarkGetDirect(b *testing.B) {
	val := "bar"
	fn := func(key string) interface{} {
		return val
	}
	benchmark(b, fn)
}

func BenchmarkGetAtomic(b *testing.B) {
	val := atomic.Value{}
	val.Store(1)
	fn := func(key string) interface{} {
		return val.Load()
	}
	benchmark(b, fn)
}

func BenchmarkGetMap(b *testing.B) {
	m := map[string]string{
		"foo": "bar",
	}
	fn := func(key string) interface{} {
		return m[key]
	}
	benchmark(b, fn)
}

func BenchmarkGetSyncMap(b *testing.B) {
	m := sync.Map{}
	m.Store("foo", "bar")
	fn := func(key string) interface{} {
		v, _ := m.Load(key)
		return v
	}
	benchmark(b, fn)
}

type TestConfig struct {
	Foo string
	Bar int64
}

func (tc *TestConfig) String() string {
	b, _ := json.Marshal(tc)
	return string(b)
}
