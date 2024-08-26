package sysutils

import (
	"runtime"
)

func KillGoRoutine() func() {
	return func() {
		runtime.Goexit()
	}

}

func GetGoRoutines() int {
	return runtime.NumGoroutine()
}
