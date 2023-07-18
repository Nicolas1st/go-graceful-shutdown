package shutdown

import "sync"

type GracefulShutter struct {
	wg       *sync.WaitGroup
	stopChan chan struct{}
}

func NewGracefulShutter() *GracefulShutter {
	return &GracefulShutter{
		wg:       &sync.WaitGroup{},
		stopChan: make(chan struct{}),
	}
}

func (s *GracefulShutter) RegOp() error {
	select {
	case <-s.stopChan:
		return ErrFinishedRegistration
	default:
		s.wg.Add(1)
		return nil
	}
}

func (s *GracefulShutter) UnregOp() {
	s.wg.Done()
}

func (s *GracefulShutter) Shutdown() {
	close(s.stopChan)
	s.wg.Wait()
}
