package main

import (
	"fmt"
	"sync"
)

type Obj struct {
	//同步锁
	mu  sync.Mutex
	val string
}

type Obj2 struct {
	//读写锁
	rwLock sync.RWMutex
	val    string
}

func main() {
	var o Obj
	//加锁
	o.mu.Lock()
	o.val = "aaa"
	o.mu.Unlock()

	fmt.Printf("o.val = %s\n", o.val)

	var o2 Obj2
	//读锁
	o2.rwLock.RLock()
	fmt.Printf("o2.val = %s\n", o2.val)
	o2.rwLock.RUnlock()

	//写锁
	o2.rwLock.Lock()
	o2.val = "bbb"
	fmt.Printf("o2.val = %s\n", o2.val)
	o2.rwLock.Unlock()

	//确保函数只被调用一次
	var caller sync.Once
	caller.Do(func() { callMe(100) })
	caller.Do(func() { callMe(200) }) //不会执行, caller.Do调用一次后,再调用就不起作用

	for i := 0; i < 10; i++ {
		//只会调用一次
		//go caller.Do(func() { callMe(1) })
		go callMe(i)
	}

	for {

	}
}

func callMe(i int) {
	fmt.Printf("callMe id=%d\n", i)
}
