package timetest

import (
	"fmt"
	"testing"
	"time"
)

// ticker 即断续器，和钟表一样每隔一段时间指针往前走一点

func TestCase1(_ *testing.T) {
	ticker := time.NewTicker(time.Second)
	for tt := range ticker.C {
		fmt.Printf("tt: %v\n", tt)
	}
}

func TestCase2(_ *testing.T) {
	ticker := time.NewTicker(time.Second * 2)
	for {
		select {
		case tt := <-ticker.C:
			fmt.Printf("tt: %v\n", tt)
		case tt := <-time.After(time.Millisecond * 3500):
			fmt.Printf("tt: %v\n", tt)
			ticker.Stop()
			return

		}
	}

}
