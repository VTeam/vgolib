package traverseStruct

import (
	"fmt"
	"reflect"
	"testing"
)

type Ref struct {
	school string `json:"school"`
}
type A struct {
	name     string `json:"name,omitempty"    xml:"name,omitempty"`
	age      int    `json:"age"`
	ref      *Ref   `json:"ref"`
	firends  []*A   `json:"firends"`
	interest map[string]string
}

func TestDFSBFSTraverseStruct(t *testing.T) {

	a := A{name: "zs", age: 1}
	tt := reflect.TypeOf(a)
	// BFSTraverseStruct(t)
	fmt.Printf("{")
	DFSTraverseStruct(tt)
	fmt.Printf("}\n")

}
