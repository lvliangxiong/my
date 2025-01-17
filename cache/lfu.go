package cache

import (
	"container/list"
	"math"
)

type entryWithFreq struct {
	entry
	freq int
}

// Try this implementation in:
// https://leetcode.cn/problems/lfu-cache/
type LFU struct {
	cap     int // >= 0, 0 indicates unlimited.
	minFreq int

	m        map[any]*list.Element
	freqList map[int]*list.List
}

func NewLFU(cap int) *LFU {
	if cap < 0 {
		panic("negative capacity")
	}
	return &LFU{
		cap:      cap,
		minFreq:  math.MaxInt,
		m:        make(map[any]*list.Element),
		freqList: make(map[int]*list.List),
	}
}

func (c *LFU) Get(key any) any {
	if _, ok := c.m[key]; ok {
		// Hit cache. Increase the freq.
		newEle := c.incrFreq(key)
		return newEle.Value.(*entryWithFreq).value
	}
	return nil
}

func (c *LFU) Set(key any, value any) {
	if _, ok := c.m[key]; !ok {
		// Insert a new entry.
		if c.cap > 0 && len(c.m) == c.cap {
			// When full, evict one entry first.
			evictKey := c.freqList[c.minFreq].Back().Value.(*entryWithFreq).key
			c.remove(evictKey)
		}
		c.add(key, value, 1)
		// Fix minFreq.
		c.minFreq = 1
		return
	}

	// Update existed entry.
	newEle := c.incrFreq(key)
	newEle.Value.(*entryWithFreq).value = value
}

func (c *LFU) Len() int {
	return len(c.m)
}

func (c *LFU) incrFreq(key any) *list.Element {
	removedKvf := c.remove(key)
	if removedKvf == nil {
		return nil
	}
	// Here we don't need fix minFreq.
	removedKvf.freq++
	return c.addEntryWithFreq(removedKvf)
}

// Remove may break minFreq.
// Call to this method must make sure minFreq can be fixed.
// Remove returns deleted entryWithFreq.
func (c *LFU) remove(key any) *entryWithFreq {
	ele, ok := c.m[key]
	if !ok {
		return nil
	}
	kvf := ele.Value.(*entryWithFreq)
	c.freqList[kvf.freq].Remove(ele)
	if c.freqList[kvf.freq].Len() == 0 {
		delete(c.freqList, kvf.freq)
	}
	delete(c.m, key)
	c.tryFixMinFreq() // Here we try to fix minFreq, but it's not enough.
	return kvf
}

func (c *LFU) add(key, value any, freq int) *list.Element {
	kvf := &entryWithFreq{
		entry: entry{
			key:   key,
			value: value,
		},
		freq: freq,
	}
	return c.addEntryWithFreq(kvf)
}

func (c *LFU) addEntryWithFreq(kvf *entryWithFreq) *list.Element {
	key, freq := kvf.key, kvf.freq
	if c.freqList[freq] == nil {
		c.freqList[freq] = new(list.List)
	}
	ele := c.freqList[freq].PushFront(kvf)
	c.m[key] = ele
	c.minFreq = min(c.minFreq, freq)
	return ele
}

func (c *LFU) tryFixMinFreq() {
	if c.freqList[c.minFreq] == nil {
		c.minFreq++
	}
}
