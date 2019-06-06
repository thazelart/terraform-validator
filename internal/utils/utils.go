// Package utils bring global functions
package utils

import (
	"log"
	"os"
)

// LogFatal is the log.Fatal go built-in function by default. Permit to change
// the behaviour of that variable in order to test the EnsureOrFatal function
var LogFatal = log.Fatal

// EnsureOrFatal ensures the error in nil or uses LogFatal
func EnsureOrFatal(err error) {
	if err != nil {
		LogFatal(err)
	}
}

// OkOrFatal ensures the answer was ok or fatal
func OkOrFatal(ok bool, message string) {
	if !ok {
		LogFatal(message)
	}
}

// FileExists return true is the given file exists, else false
func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
