package shutdown

import (
	"sync"
)

type GracefulShutter struct {
	opsCount           *int64
	stoppedRegistering bool
	mu                 *sync.RWMutex
}

func NewConcurrentOpsRegister() *GracefulShutter {
	return &GracefulShutter{
		opsCount:           new(int64),
		stoppedRegistering: false,
		mu:                 &sync.RWMutex{},
	}
}
