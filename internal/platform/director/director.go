package director

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Task interface {
	Start()
	Stop()
}

var wg sync.WaitGroup

func Run(tasks ...Task) {
	ctx := handleSignal()
	RunWithContext(ctx, tasks...)
}

func RunWithContext(ctx context.Context, tasks ...Task) {
	wg.Add(len(tasks))
	for _, t := range tasks {
		go func(t Task) {
			defer wg.Done()
			t.Start()
			<-ctx.Done()
			t.Stop()
		}(t)
	}
}

func Wait() {
	wg.Wait()
}

// handleSignal returns a context that gets cancelled on SIGINT or SIGTERM.
func handleSignal() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		fmt.Print("\r")
		cancel()
	}()
	return ctx
}
