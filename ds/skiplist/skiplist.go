package skiplist

import "math/rand"

const (
	MaxLevel = 32
	P        = 0.5
)

type Node struct {
	Val  int
	Next []*Node
}

// Try this implementation in:
// https://leetcode.cn/problems/design-skiplist/
type Skiplist struct {
	Dummy *Node
	Level int
}

func New() *Skiplist {
	return &Skiplist{
		Dummy: &Node{Next: make([]*Node, MaxLevel)},
		Level: 1,
	}
}

func (sl *Skiplist) Search(target int) bool {
	p := sl.Dummy
	for i := sl.Level - 1; i >= 0; i-- {
		for p.Next[i] != nil && p.Next[i].Val < target {
			p = p.Next[i]
		}
	}
	return p.Next[0] != nil && p.Next[0].Val == target
}

func (sl *Skiplist) Add(num int) {
	// Prepare the new node including height.
	h := 1
	for h < MaxLevel && rand.Float64() < P {
		h++
	}
	newNode := &Node{Val: num, Next: make([]*Node, h)}

	// Update level.
	if sl.Level < h {
		sl.Level = h
	}

	// Find the right position, insert it at each level
	// from top to bottom.
	p := sl.Dummy
	for i := sl.Level - 1; i >= 0; i-- {
		for p.Next[i] != nil && p.Next[i].Val < num {
			p = p.Next[i]
		}
		if i < h {
			newNode.Next[i] = p.Next[i]
			p.Next[i] = newNode
		}
	}
}

func (sl *Skiplist) Erase(num int) bool {
	p := sl.Dummy
	found := false
	for i := sl.Level - 1; i >= 0; i-- {
		for p.Next[i] != nil && p.Next[i].Val < num {
			p = p.Next[i]
		}
		if p.Next[i] != nil && p.Next[i].Val == num {
			p.Next[i] = p.Next[i].Next[i]
			found = true
		}
	}
	return found
}
