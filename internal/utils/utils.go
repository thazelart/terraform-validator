// utils package bring global functions
package utils

import (
	"log"
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
