package director

import (
	"context"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type SimpleTask struct {
	started bool
	stopped bool
}

func (t *SimpleTask) Start() {
	t.started = true
}

func (t *SimpleTask) Stop() {
	t.stopped = true
}

func TestRunWithContext(t *testing.T) {
	assert := assert.New(t)
	s := &SimpleTask{}
	ctx, cancel := context.WithCancel(context.Background())

	assert.False(s.started)
	assert.False(s.stopped)
	RunWithContext(ctx, s)
	time.Sleep(1 * time.Millisecond) // ensure goroutine has started
	assert.True(s.started)
	assert.False(s.stopped)
	cancel()
	Wait()
	assert.True(s.stopped)
}

func TestWaitEmpty(t *testing.T) {
	c := make(chan struct{})
	go func() {
		Wait() // should not hang
		c <- struct{}{}
	}()
	select {
	case <-c:
	case <-time.After(1 * time.Millisecond):
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
	case <-time.After(2 * time.Millisecond):
		t.Fatal("Wait() is taking too long")
	}
}

func TestRunSIGINT(t *testing.T) {
	assert := assert.New(t)
	s := &SimpleTask{}
	Run(s)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	time.Sleep(1 * time.Millisecond)
	assert.True(s.stopped)
}

func TestRunSIGTERM(t *testing.T) {
	assert := assert.New(t)
	s := &SimpleTask{}
	Run(s)
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	time.Sleep(1 * time.Millisecond)
	assert.True(s.stopped)
}
