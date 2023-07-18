package shutdown

import (
	"sync/atomic"
)

type GracefulShutter struct {
	opsCount *int64
	stopChan chan struct{}
}

func NewGracefulShutter() *GracefulShutter {
	return &GracefulShutter{
		opsCount: new(int64),
		stopChan: make(chan struct{}),
	}
}

func (s *GracefulShutter) RegOp() error {
	select {
	case <-s.stopChan:
		return ErrRegistrationStopped
	default:
		atomic.AddInt64(s.opsCount, 1)
		return nil
	}
}

func (s *GracefulShutter) UnregOp() error {
	// catch unintended lib use
	if atomic.LoadInt64(s.opsCount) == 0 {
		return ErrNothingToRegister
	}

	// catch unintended lib use
	if newOpsCount := atomic.AddInt64(s.opsCount, -1); newOpsCount < 0 {
		return ErrNegativeOpsCount
	}

	return nil
}

func (s *GracefulShutter) ShutdownGracefully() (err error) {
	defer func() {
		if recover() == nil {
			err = nil
		} else {
			err = ErrRepetitiveStopping
		}
	}()

	close(s.stopChan)

	return err
}
