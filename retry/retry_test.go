package retry

import (
	"fmt"
	"testing"
	"time"
)

var (
	called1   int
	called2   int
	hasCalled bool
)

func targetFuncAllErr() error {
	// do something
	fmt.Println("do something", called1)
	called1++
	return ErrConflict
}

func targetFunc3Err() error {
	// do something
	fmt.Println("do something", called2)
	called2++
	if called2 == 3 {
		return nil
	}
	return ErrConflict
}

func targetFuncNoErr() error {
	// do something
	fmt.Println("do something")
	hasCalled = true
	return nil
}

func TestRetryOnConflict(t *testing.T) {

	var TestBackoff1 = Backoff{
		Steps:    5,
		Duration: 10 * time.Millisecond,
		Factor:   1.0,
		Jitter:   0.1,
	}

	// case1: 5 Steps all error
	err := RetryOnConflict(TestBackoff1, targetFuncAllErr)
	if err != ErrConflict {
		t.Errorf("err: %s", err)
	}
	if called1 != TestBackoff1.Steps {
		t.Error("called less")
	}
	// case2: 3 Steps error
	var TestBackoff2 = Backoff{
		Steps:    5,
		Duration: 10 * time.Millisecond,
		Factor:   1.0,
		Jitter:   0.1,
	}
	err = RetryOnConflict(TestBackoff2, targetFunc3Err)
	if err != nil {
		t.Errorf("err: %s", err)
	}
	if called2 != 3 {
		t.Error("called times error")
	}

	// case3: no error
	var TestBackoff3 = Backoff{
		Steps:    5,
		Duration: 10 * time.Millisecond,
		Factor:   1.0,
		Jitter:   0.1,
	}
	err = RetryOnConflict(TestBackoff3, targetFuncNoErr)
	if err != nil {
		t.Errorf("err: %s", err)
	}
	if !hasCalled {
		t.Error("no called")
	}
}
