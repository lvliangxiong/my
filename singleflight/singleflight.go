package singleflight

import (
	"fmt"
	"sync"
)

type Group struct {
	callMap     map[string]*call
	callMapLock sync.Mutex
}

type call struct {
	cnt  int
	done chan struct{}

	ret interface{}
	err error
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
	g.callMapLock.Lock()

	// 延迟初始化
	if g.callMap == nil {
		g.callMap = make(map[string]*call, 64)
	}

	if c, ok := g.callMap[key]; ok {
		c.cnt++
		g.callMapLock.Unlock()

		// 等待 call 执行完成
		<-c.done
		return c.ret, c.err, true
	}

	c := new(call)
	c.done = make(chan struct{})
	g.callMap[key] = c
	g.callMapLock.Unlock()

	// 实际执行 fn
	func() {
		defer func() {
			if r := recover(); r != nil {
				c.err = fmt.Errorf("%v", r)
			}
		}()

		c.ret, c.err = fn()
	}()

	g.callMapLock.Lock()
	delete(g.callMap, key)
	close(c.done) // 通知其它等待结果的调用者：结果已经 ready
	g.callMapLock.Unlock()

	return c.ret, c.err, c.cnt > 0
}
