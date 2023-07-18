package shutdown

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCorrectConcurrencyUse(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			t.Errorf("Error, Panic when using correctly")
		} else {
			t.Log("No panic, as it should be when the lib is being used correctly")
		}
	}()

	s := NewGracefulShutter()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		if err := s.ShutdownGracefully(); err != nil {
			t.Error("Error, graceful shutdown did not work properly when the lib is used as inteded")
		} else {
			t.Logf("Graceful shutdown stopped successfully, ops count is %d", *s.opsCount)
		}
	}()

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			s.RegOp()
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			s.UnregOp()

		}()
	}

	wg.Wait()
}
