package skiplist

import "math/rand"

const (
	maxLevel int     = 16 // Should be enough for 2^16 elements
	p        float32 = 0.25
)

// Element is an Element of a skiplist.
type Element struct {
	Score   float64
	Value   interface{}
	forward []*Element
}

// Next returns first element after e.
func (e *Element) Next() *Element {
	if e != nil {
		return e.forward[0]
	}
	return nil
}

func newElement(score float64, value interface{}, level int) *Element {
	return &Element{
		Score:   score,
		Value:   value,
		forward: make([]*Element, level),
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
	return sl.header.forward[0]
}

// Search the skiplist to findout element with the given score.
// Returns (*Element, true) if the given score present, otherwise returns (nil, false).
func (sl *SkipList) Search(score float64) (element *Element, ok bool) {
	x := sl.header
	for i := sl.level - 1; i >= 0; i-- {
		for x.forward[i] != nil && x.forward[i].Score < score {
			x = x.forward[i]
		}
	}
	x = x.forward[0]
	if x != nil && x.Score == score {
		return x, true
	}
	return nil, false
}

// New returns a new empty SkipList.
func New() *SkipList {
	return &SkipList{
		header: &Element{forward: make([]*Element, maxLevel)},
	}
}

func randomLevel() int {
	level := 1
	for rand.Float32() < p && level < maxLevel {
		level++
	}
	return level
}
