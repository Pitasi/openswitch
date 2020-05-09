package director

import (
	"context"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SimpleTask struct {
	sync.Mutex
	started bool
	stopped bool
}

func (t *SimpleTask) Start() {
	t.Lock()
	defer t.Unlock()
	t.started = true
}

func (t *SimpleTask) Stop() {
	t.Lock()
	defer t.Unlock()
	t.stopped = true
}

func (t *SimpleTask) hasStarted() bool {
	t.Lock()
	defer t.Unlock()
	return t.started
}

func (t *SimpleTask) hasStopped() bool {
	t.Lock()
	defer t.Unlock()
	return t.stopped
}

func TestRunWithContext(t *testing.T) {
	assert := assert.New(t)
	s := &SimpleTask{}
	ctx, cancel := context.WithCancel(context.Background())

	assert.False(s.hasStarted())
	assert.False(s.hasStopped())
	RunWithContext(ctx, s)
	time.Sleep(10 * time.Millisecond) // ensure goroutine has started
	assert.True(s.hasStarted())
	assert.False(s.hasStopped())
	cancel()
	Wait()
	assert.True(s.hasStopped())
}

func TestWaitEmpty(t *testing.T) {
	c := make(chan struct{})
	go func() {
		Wait() // should not hang
		c <- struct{}{}
	}()
	select {
	case <-c:
	case <-time.After(10 * time.Millisecond):
		t.Fatal("Wait() did not return immediately")
	}
}

func TestWait(t *testing.T) {
	s := &SimpleTask{}
	ctx, cancel := context.WithCancel(context.Background())
	RunWithContext(ctx, s)

	// cancel context after 1ms
	go func() {
		time.Sleep(1 * time.Millisecond)
		cancel()
	}()

	c := make(chan struct{})
	// Wait should halt execution for 1ms
	go func() {
		Wait()
		c <- struct{}{}
	}()

	// check whether Wait() is taking more than 2ms
	select {
	case <-c:
	case <-time.After(5 * time.Millisecond):
		t.Fatal("Wait() is taking too long")
	}
}

func TestRunSIGINT(t *testing.T) {
	assert := assert.New(t)
	s := &SimpleTask{}
	Run(s)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	time.Sleep(10 * time.Millisecond)
	assert.True(s.hasStopped())
}

func TestRunSIGTERM(t *testing.T) {
	assert := assert.New(t)
	s := &SimpleTask{}
	Run(s)
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	time.Sleep(10 * time.Millisecond)
	assert.True(s.hasStopped())
}
