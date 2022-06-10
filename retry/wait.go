package retry

import (
	"context"
	"errors"
	"math/rand"
	"time"
)

func Jitter(duration time.Duration, maxFactor float64) time.Duration {
	if maxFactor <= 0.0 {
		maxFactor = 1.0
	}
	wait := duration + time.Duration(rand.Float64()*maxFactor*float64(duration))
	return wait
}

var ErrWaitTimeout = errors.New("timed out waiting for the condition")

type ConditionFunc func() (done bool, err error)

type ConditionWithContextFunc func(context.Context) (done bool, err error)

func (cf ConditionFunc) WithContext() ConditionWithContextFunc {
	return func(context.Context) (done bool, err error) {
		return cf()
	}
}

func runConditionWithCrashProtection(condition ConditionFunc) (bool, error) {
	return runConditionWithCrashProtectionWithContext(context.TODO(), condition.WithContext())
}

func runConditionWithCrashProtectionWithContext(ctx context.Context, condition ConditionWithContextFunc) (bool, error) {
	// TODO: condition may be panic
	return condition(ctx)
}

type Backoff struct {
	Duration time.Duration

	Factor float64

	Jitter float64

	Steps int

	Cap time.Duration
}

func (b *Backoff) Step() time.Duration {
	if b.Steps == 0 {
		if b.Jitter > 0 {
			return Jitter(b.Duration, b.Jitter)
		}
		return b.Duration
	}
	b.Steps--

	duration := b.Duration

	if b.Factor != 0 {
		b.Duration = time.Duration(float64(b.Duration) * b.Factor)
		if b.Cap > 0 && b.Duration > b.Cap {
			b.Duration = b.Cap
			b.Steps = 0
		}
	}

	if b.Jitter > 0 {
		duration = Jitter(duration, b.Jitter)
	}
	return duration
}

func ExponentialBackoff(backoff Backoff, condition ConditionFunc) error {
	for backoff.Steps > 0 {
		if ok, err := runConditionWithCrashProtection(condition); err != nil || ok {
			return err
		}
		if backoff.Steps == 1 {
			break
		}
		time.Sleep(backoff.Step())
	}
	return ErrWaitTimeout
}
