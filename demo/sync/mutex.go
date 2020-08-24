package main

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

type RWMutex struct {
	sync.RWMutex
}

type m struct {
	w           sync.Mutex
	writeSem    int32
	readerSem   int32
	readerCount int32
	readerWait  int32
}

func (rw *RWMutex) ReaderCount() int {
	v := (*m)(unsafe.Pointer(&rw.RWMutex))
	c := int(v.readerCount)
	if c < 0 {
		c = int(v.readerWait)
	}
	return c
}

func (rw *RWMutex) WriterCount() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&rw.RWMutex)))
	v = v >> mutexWaiterShift
	v = v + (v & mutexLocked)
	return int(v)
}
