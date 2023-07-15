package shutdown

import (
	"sync/atomic"
)

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
	if atomic.LoadInt64(s.opsCount) == 0 {
		return ErrNothingToRegister
	}

	atomic.AddInt64(s.opsCount, -1)

	s.mu.RLock()
	if s.stoppedRegistering && atomic.LoadInt64(s.opsCount) == 0 {
		s.finishedWorkingChan <- struct{}{}
	}
	s.mu.RUnlock()

	return nil
}

func (s *GracefulShutter) StopRegistering() {
	s.mu.Lock()
	s.stoppedRegistering = true
	s.mu.Unlock()

	<-s.finishedWorkingChan
}
