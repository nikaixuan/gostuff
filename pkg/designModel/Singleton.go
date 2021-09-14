package designModel

import (
	"sync"
)

type tool struct{}

var instance1 *tool
var instance2 = new(tool)
var instance3 *tool
var lock sync.Mutex
var once sync.Once

func newToolLazy() *tool {
	lock.Lock()
	defer lock.Unlock()
	if instance1 == nil {
		instance1 = new(tool)
	}
	return instance1
}

func newToolHungry() *tool {
	return instance2
}

func newToolDoubleCheck() *tool {
	if instance3 == nil {
		lock.Lock()
		if instance3 == nil {
			instance3 = new(tool)
		}
		lock.Unlock()
	}
	return instance3
}

// Also double check
func newToolOnce() *tool {
	once.Do(func() {
		instance3 = new(tool)
	})
	return instance3
}
