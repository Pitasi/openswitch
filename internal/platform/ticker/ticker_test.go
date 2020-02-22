package ticker

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type counter int64

func (c *counter) inc(t time.Time) {
	(*c)++
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
			assert.GreaterOrEqual(int64(c), expectedCount-1)
			assert.LessOrEqual(int64(c), expectedCount+1)
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

	assert.Equal(1, int(c))
}
