package shutdown

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCorrectUse(t *testing.T) {
	s := NewGracefulShutter()

	wg := &sync.WaitGroup{}

	shutdownChan := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()

		select {
		case <-time.After(1500 * time.Millisecond):
			t.Error("Must have finished by now, most likely there is a deadlock")
		case <-shutdownChan:
			t.Log("Success, it did shutdown")
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		// waiting for an app to run for some time before shutting down
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

		s.Shutdown()
		close(shutdownChan)
	}()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			if err := s.RegOp(); err == nil {
				defer s.UnregOp()
			}

			// doing some work
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}()
	}

	wg.Wait()
}
