package shutdown

import (
	"testing"
	"time"
)

func TestShutdownGracefully(t *testing.T) {
	s := NewGracefulShutter()

	select {
	case <-s.stopChan:
		t.Error("Error, The channel must be open")
	case <-time.After(100 * time.Millisecond):
		t.Log("The channel is open in the begging as it is to be expected")
	}

	if err := s.ShutdownGracefully(); err != nil {
		t.Error("Error, There should be no error when calling the method for the first time")
	} else {
		t.Log("The method returns no error as it is to be expected")
	}

	select {
	case <-s.stopChan:
		t.Log("The channel is open in the begging as it is to be expected")
	case <-time.After(100 * time.Millisecond):
		t.Error("Error, The channel must be closed")
	}

	if err := s.ShutdownGracefully(); err != nil {
		t.Log("There is an error when calling the method more than once")
	} else {
		t.Error("Error, The method returns no error")
	}

	select {
	case <-s.stopChan:
		t.Log("The channel is still closed, as it should be")
	case <-time.After(100 * time.Millisecond):
		t.Error("Error, The channel must still be closed")
	}
}
