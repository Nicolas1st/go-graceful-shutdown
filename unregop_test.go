package shutdown

import "testing"

func TestUnregOp(t *testing.T) {
	s := NewGracefulShutter()

	if err := s.UnregOp(); err != nil {
		t.Log("Returning an error, when there is nothing to unregister, as expected")
	} else {
		t.Error("Error, Not returning an error, where there is nothing to unregister")
	}

	s.RegOp()
	if err := s.UnregOp(); err != nil {
		t.Error("Error, Returning an error, where there is something to unregister")
	} else {
		t.Log("Not returning an error, when there is something to unregister, as expected")
	}

	// checking for possible data races, can be the result of forgeting to call a RegOp method
	*s.opsCount = -1
	if err := s.UnregOp(); err != nil {
		t.Log("Returning an error, when the result is a negative number, as expected")
	} else {
		t.Error("Error, not returning an error, where the result is a negative number")
	}
}
