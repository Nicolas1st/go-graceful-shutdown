package shutdown

import "errors"

var (
	ErrRegistrationStopped = errors.New("operations can no longer be registered")
	ErrNothingToRegister   = errors.New("the op count is already 0, nothing to unregister")
	ErrNegativeOpsCount    = errors.New("negative ops count")
	ErrRepetitiveStopping  = errors.New("a shutter can not be stopped again after being stopped once")
)
