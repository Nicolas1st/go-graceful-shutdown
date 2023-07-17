package shutdown

import "testing"

func TestRegOp(t *testing.T) {
	s := NewGracefulShutter()

	if err := s.RegOp(); err != nil {
		t.Error(err.Error())
	} else {
		t.Log("The RegOp op did not incure any errors")
	}

	if *s.opsCount == 1 {
		t.Log("The ops count is to 1, as it should be after registering an op")
	} else {
		t.Errorf("The ops count is equal to %d, though it shoudl be 1, after registering only 1 op", *s.opsCount)
	}

	for i := 0; i < 10; i++ {
		if err := s.RegOp(); err != nil {
			t.Error(err.Error())
		} else {
			t.Log("The RegOp op did not incure any errors")
		}
	}

	if *s.opsCount == 11 {
		t.Log("Registed everything")
	} else {
		t.Errorf("Wrong number of registered ops %d", *s.opsCount)
	}

	s.ShutdownGracefully()
	err := s.RegOp()
	if err != nil {
		t.Log("Returned an error after stoping the shutter, as expected")
	} else {
		t.Errorf("Did not return an error after the shutter being closed, which is not an expected behvior")
	}

	if *s.opsCount != 11 {
		t.Errorf("The opsCount has a wrong value after the last RegOp call %d, though it must 11", *s.opsCount)
	} else {
		t.Log("The ops count is still equal to 11, as expected")
	}
}
