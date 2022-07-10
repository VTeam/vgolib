package skiplist

import (
	"fmt"
	"math/rand"
)

const (
	maxLevel int     = 16 // Should be enough for 2^16 elements
	p        float32 = 0.25
)

// Element is an Element of a skiplist.
type Element struct {
	Score float64
	Value interface{}
	next  []*Element
}

// Next returns first element after e.
func (e *Element) Next() *Element {
	if e != nil {
		return e.next[0]
	}
	return nil
}

func newElement(score float64, value interface{}, level int) *Element {
	return &Element{
		Score: score,
		Value: value,
		next:  make([]*Element, level),
	}
}

// SkipList represents a skiplist.
// The zero value from SkipList is an empty skiplist ready to use.
type SkipList struct {
	header *Element // header is a dummy element
	len    int      // current skiplist length，header not included
	level  int      // current skiplist level，header not included
}

// Front returns first element in the skiplist which maybe nil.
func (sl *SkipList) Front() *Element {
	return sl.header.next[0]
}

// Search the skiplist to findout element with the given score.
// Returns (*Element, true) if the given score present, otherwise returns (nil, false).
func (sl *SkipList) Search(score float64) (element *Element, ok bool) {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].Score < score {
			x = x.next[i]
		}
	}
	x = x.next[0]
	if x != nil && x.Score == score {
		return x, true
	}
	return nil, false
}

// Insert (score, value) pair to the skiplist and returns pointer of element.
func (sl *SkipList) Insert(score float64, value interface{}) *Element {
	update := make([]*Element, maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].Score < score {
			x = x.next[i]
		}
		update[i] = x
	}
	x = x.next[0]

	// Score already presents, replace with new value then return
	if x != nil && x.Score == score {
		x.Value = value
		return x
	}

	level := randomLevel()
	if level > sl.level {
		level = sl.level + 1
		update[sl.level] = sl.header
		sl.level = level
	}
	e := newElement(score, value, level)
	for i := 0; i < level; i++ {
		e.next[i] = update[i].next[i]
		update[i].next[i] = e
	}
	sl.len++
	return e
}

func (sl *SkipList) Len() int {

	x := sl.header
	var r int
	for x != nil {
		x = x.next[0]
		r++
	}
	return r - 1

}

// Delete remove and return element with given score, return nil if element not present
func (sl *SkipList) Delete(score float64) *Element {
	update := make([]*Element, maxLevel)
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.next[i] != nil && x.next[i].Score < score {
			x = x.next[i]
		}
		update[i] = x
	}
	x = x.next[0]

	if x != nil && x.Score == score {
		for i := 0; i < sl.level; i++ {
			if update[i].next[i] != x {
				return nil
			}
			update[i].next[i] = x.next[i]
		}
		sl.len--
	}
	return x
}

func (sl *SkipList) Traverse() {
	for e := sl.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// New returns a new empty SkipList.
func New() *SkipList {
	return &SkipList{
		header: &Element{next: make([]*Element, maxLevel)},
	}
}

func randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}
