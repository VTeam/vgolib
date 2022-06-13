package signals

import (
	"context"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

func TestSetupSignalHandler(t *testing.T) {
	ctx := SetupSignalHandler()

	var (
		done int32
		wait chan struct{}
	)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// do something cleanup
				t.Log("done")
				atomic.StoreInt32(&done, 1)
				return
			case <-wait:
				t.Log("wait")
			}
		}

	}(ctx)

	t.Logf("sighup...")
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	time.Sleep(100 * time.Millisecond)
	if atomic.LoadInt32(&done) != 1 {
		t.Error("singal not capture")
	}

}
