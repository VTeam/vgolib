package timetest

import (
	"fmt"
	"testing"
	"time"
)

//timer即计时器，和iphone的计时器一样，到时间后提醒

func TestCase3(_ *testing.T) {
	timer := time.NewTimer(time.Second * 1)
	fmt.Printf("time.Now(): %v\n", time.Now())
	tt := <-timer.C
	fmt.Printf("tt: %v\n", tt)
}

func TestCase4(_ *testing.T) {
	for {
		select {
		case tt := <-time.After(time.Second):
			fmt.Printf("tt: %v\n", tt)
		case yy := <-time.NewTimer(time.Second).C:
			fmt.Printf("yy: %v\n", yy)
		}
	}
}

func TestCase5(_ *testing.T) {
	for {
		timer := time.NewTimer(time.Second * 2)
		select {
		case tt := <-timer.C:
			fmt.Printf("tt: %v\n", tt)
		// there has some case to change timer
		case <-time.After(time.Second * 1):
			timer.Stop()
		}

	}

}

func TestCase7(_ *testing.T) {
	t1 := time.Now()
	time.Sleep(time.Second)
	defer func() {
		fmt.Printf("time.Since(t1): %v\n", time.Since(t1))
		fmt.Printf("time.Now().Sub(t1): %v\n", time.Now().Sub(t1))
	}()
	fmt.Println(t1.Clock())
	fmt.Printf("time.Until(t1): %v\n", time.Until(t1))
	fmt.Printf("time.Now().YearDay(): %v\n", time.Now().YearDay())
}

const (
	TIMEFMT = "2006-01-02 15:04:05"
)

func TestTimeFormat(_ *testing.T) {
	fmt.Printf("time.Now().Format(TIMEFMT): %v\n", time.Now().Format(TIMEFMT))
	tt, _ := time.Parse(TIMEFMT, "2022-09-18 20:04:20")
	fmt.Printf("tt.Add(time.Hour).Format(TIMEFMT): %v\n", tt.Add(time.Hour).Format(TIMEFMT))
}
