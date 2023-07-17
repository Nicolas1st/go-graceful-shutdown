package shutdown

import (
	"sync"
	"sync/atomic"
)

type GracefulShutter struct {
	opsCount            *int64
	stoppedRegistering  bool
	mu                  *sync.RWMutex
	finishedWorkingChan chan struct{}
}

func NewConcurrentOpsRegister() *GracefulShutter {
	return &GracefulShutter{
		opsCount:            new(int64),
		stoppedRegistering:  false,
		mu:                  &sync.RWMutex{},
		finishedWorkingChan: make(chan struct{}),
	}
}

func (s *GracefulShutter) RegisterOp() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.stoppedRegistering {
		return ErrRegistrationStopped
	}

	atomic.AddInt64(s.opsCount, 1)

	return nil
}

func (s *GracefulShutter) UnregisterOp() error {
	if atomic.LoadInt64(s.opsCount) != 0 {
		atomic.AddInt64(s.opsCount, -1)
		return nil
	}

	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.stoppedRegistering && atomic.LoadInt64(s.opsCount) == 0 {
		s.finishedWorkingChan <- struct{}{}
		return nil
	} else {
		return ErrNothingToRegister
	}
}

func (s *GracefulShutter) StopRegistering() {
	s.mu.Lock()
	s.stoppedRegistering = true
	s.mu.Unlock()

	<-s.finishedWorkingChan
}
