package retry

import (
	"errors"
	"time"
)

var DefaultRetry = Backoff{
	Steps:    5,
	Duration: 10 * time.Millisecond,
	Factor:   1.0,
	Jitter:   0.1,
}

var DefaultBackoff = Backoff{
	Steps:    4,
	Duration: 10 * time.Millisecond,
	Factor:   5.0,
	Jitter:   0.1,
}

var ErrConflict = errors.New("Conflict")

func IsConflict(err error) bool {
	return err == ErrConflict
}

func OnError(backoff Backoff, retriable func(error) bool, fn func() error) error {
	var lastErr error
	err := ExponentialBackoff(backoff, func() (bool, error) {
		err := fn()
		switch {
		// fn not error, return direct
		case err == nil:
			return true, nil
		// fn error with define, retry
		case retriable(err):
			lastErr = err
			return false, nil
		// fn error with other error, return direct, when retry timeout, case there
		default:
			return false, err
		}
	})
	if err == ErrWaitTimeout {
		err = lastErr
	}
	return err
}

func RetryOnConflict(backoff Backoff, fn func() error) error {
	return OnError(backoff, IsConflict, fn)
}
