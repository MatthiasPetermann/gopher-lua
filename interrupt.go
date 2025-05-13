package lua

import "sync/atomic"

// abortFlag holds the global abort signal state.
// It should be accessed via SetAbort, ResetAbort, and ShouldAbort.
var abortFlag atomic.Bool

// SetAbort sets the abort flag to true.
// This is typically called from an external signal handler
// to cooperatively interrupt Lua execution.
func SetAbort() {
	abortFlag.Store(true)
}

// ResetAbort clears the abort flag.
// This should be called before starting a new Lua execution,
// to ensure the abort state is clean.
func ResetAbort() {
	abortFlag.Store(false)
}

// ShouldAbort returns true if an abort has been requested.
// This is intended to be checked periodically within the VM loop.
func ShouldAbort() bool {
	return abortFlag.Load()
}
