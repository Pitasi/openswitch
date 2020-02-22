// Package ticker is a wrapper over time.Ticker to provide some higher level
// functions.
package ticker

import (
	"context"
	"time"
)

// PeriodicTask runs task everytime interval is passed. The task is also run
// once right after the calling of this function. To stop a PeriodicTask,
// passed context should be cancelled.
func PeriodicTask(ctx context.Context, task func(time.Time), interval time.Duration) {
	go task(time.Now())

	ticker := time.NewTicker(interval)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				task(t)
			}
		}
	}()
}
