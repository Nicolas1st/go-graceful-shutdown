package shutdown

import (
	"testing"
)

func TestIncorrectUse(t *testing.T) {
	defer func() {
		if recoverVal := recover(); recoverVal == nil {
			t.Error("No panic occured, the impl is broken now")
		} else {
			t.Logf("Incorrect lib use caused a panic to occur, as expected, %s", recoverVal)
		}
	}()

	s := NewGracefulShutter()

	s.Shutdown()

	s.RegOp()
	s.UnregOp()
}
