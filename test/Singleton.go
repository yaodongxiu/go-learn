package main

import "sync"

type Tool struct {
	count int
}

var singleton *Tool

var lock sync.Mutex
var once sync.Once

func GetInstance() *Tool {
	once.Do(func() {
		singleton = new(Tool)
		singleton.count = 1
	})
	return singleton
}

func GetInstanceByLock() *Tool {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			singleton = new(Tool)
		}
	}
	return singleton
}
