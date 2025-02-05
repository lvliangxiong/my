package cache

import (
	"container/list"
)

type entry struct {
	key, value any
}

// Try this implementation in
// https://leetcode.cn/problems/lru-cache.
type LRU struct {
	cap int

	l list.List
	m map[any]*list.Element
}

func NewLRU(cap int) *LRU {
	if cap < 0 {
		panic("negative capacity")
	}
	return &LRU{
		m:   make(map[any]*list.Element),
		cap: cap,
	}
}

func (c *LRU) Get(key any) any {
	if ele, ok := c.m[key]; ok {
		c.l.MoveToFront(ele)
		return ele.Value.(*entry).value
	}
	return nil
}

func (c *LRU) Set(key any, value any) {
	if ele, ok := c.m[key]; ok {
		// Update exist entry.
		ele.Value.(*entry).value = value
		c.l.MoveToFront(ele)
		return
	}
	// Add new entry.
	if c.cap > 0 && c.l.Len() == c.cap {
		// When full, evict one entry first.
		ele := c.l.Back()
		c.l.Remove(ele)
		delete(c.m, ele.Value.(*entry).key)
	}
	ele := c.l.PushFront(&entry{key, value})
	c.m[key] = ele
}

func (c *LRU) Len() int {
	return c.l.Len()
}
