package ticker

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type counter struct {
	sync.Mutex
	val int64
}

func (c *counter) inc(t time.Time) {
	c.Lock()
	defer c.Unlock()
	(c.val)++
}

func (c *counter) get() int64 {
	c.Lock()
	defer c.Unlock()
	return c.val
}

func TestPeriodicTask(t *testing.T) {
	tests := []struct {
		name     string
		duration time.Duration
		interval time.Duration
	}{
		{
			name:     "slow ticker",
			duration: 100 * time.Millisecond,
			interval: 75 * time.Millisecond,
		},
		{
			name:     "fast ticker",
			duration: 100 * time.Millisecond,
			interval: 1 * time.Millisecond,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)
			var c counter
			expectedCount := test.duration.Microseconds() / test.interval.Microseconds()
			ctx, cancel := context.WithCancel(context.Background())

			PeriodicTask(ctx, c.inc, test.interval)
			time.Sleep(test.duration)
			cancel()

			// we assume that count can be +/-1 because of many factors
			assert.GreaterOrEqual(c.get(), expectedCount-1)
			assert.LessOrEqual(c.get(), expectedCount+1)
		})
	}

}

func TestRunAtLeastOnce(t *testing.T) {
	assert := assert.New(t)
	var c counter
	ctx, cancel := context.WithCancel(context.Background())

	PeriodicTask(ctx, c.inc, 10*time.Hour)
	time.Sleep(1 * time.Millisecond) // ensure goroutine run
	cancel()

	assert.EqualValues(1, c.get())
}
