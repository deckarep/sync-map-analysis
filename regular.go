package main

import (
	"sync"
)

type RegularStringMap struct {
	sync.RWMutex
	internal map[string]string
}

func NewRegularStringMap() *RegularStringMap {
	return &RegularStringMap{
		internal: make(map[string]string),
	}
}

func (rm *RegularStringMap) Load(key string) (value string, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *RegularStringMap) Delete(key string) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *RegularStringMap) Store(key, value string) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

type RegularIntMap struct {
	sync.RWMutex
	internal map[int]int
}

func NewRegularIntMap() *RegularIntMap {
	return &RegularIntMap{
		internal: make(map[int]int),
	}
}

func (rm *RegularIntMap) Load(key int) (value int, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *RegularIntMap) Delete(key int) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *RegularIntMap) Store(key, value int) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}
