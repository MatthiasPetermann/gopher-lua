package lua

import "sync/atomic"

var abortFlag atomic.Bool

func SetAbort() {
	abortFlag.Store(true)
}

func ResetAbort() {
	abortFlag.Store(false)
}

func ShouldAbort() bool {
	return abortFlag.Load()
}
