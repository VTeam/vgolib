package skiplist

import (
	"fmt"
	"testing"
)

func TestSKL(*testing.T) {

	sl := New()

	sl.Insert(float64(100), "foo")

	e, ok := sl.Search(float64(100))
	fmt.Println(ok)
	fmt.Println(e.Value)
	e, ok = sl.Search(float64(200))
	fmt.Println(ok)
	fmt.Println(e)

	sl.Insert(float64(20.5), "bar")
	sl.Insert(float64(50), "spam")
	sl.Insert(float64(20), 42)

	fmt.Println(sl.Len())
	e = sl.Delete(float64(50))
	fmt.Println(e.Value)
	fmt.Println(sl.Len())

}

func TestTraverse(*testing.T) {
	sl := New()

	sl.Insert(float64(100), "foo")
	sl.Insert(float64(20.5), "bar")
	sl.Insert(float64(50), "spam")
	sl.Insert(float64(20), 42)

	sl.Traverse()
}
