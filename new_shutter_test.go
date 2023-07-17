package shutdown

import "testing"

func TestNewGracefulShutter(t *testing.T) {
	s := NewGracefulShutter()

	switch *s.opsCount {
	case 0:
		t.Log("The ops count is set to 0, as it should be")
	default:
		t.Errorf("The ops count is equal to %d when doing the init", *s.opsCount)
	}
}
