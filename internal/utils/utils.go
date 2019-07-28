// Package utils bring global functions
package utils

import (
	"log"
	"os"
	"os/exec"
	"bytes"
)

// LogFatal is the log.Fatal go built-in function by default. Permit to change
// the behaviour of that variable in order to test the EnsureOrFatal function
var LogFatal = log.Fatal

// LogFatalf is the log.Fatalf go built-in function by default. Permit to change
// the behaviour of that variable in order to test EnsureProgramInstalled function
var LogFatalf = log.Fatalf

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

// Contains check if a list contains a given string
func Contains(list []string, string string) bool {
	for _, elem := range list {
		if elem == string {
			return true
		}
	}
	return false
}

func RunSystemCommand(name string, arg ...string) (string, bool) {
	cmd := exec.Command(name, arg...)
    var stdout bytes.Buffer
    cmd.Stdout = &stdout
		cmd.Stderr = &stdout
    err := cmd.Run()
    if err != nil {
        LogFatalf("cmd.Run() failed with %s\n", err)
				return "", false
    }
    outStr := string(stdout.Bytes())
		return outStr, true
}

// EnsureProgramInstalled check if the given program is installed in the sytem.
// For example, check if terraform is installed. If not, will crash the program.
func EnsureProgramInstalled(programName string) bool {
    _, err := exec.LookPath(programName)
    if err != nil {
        LogFatalf("FATAL: %s is not installed\n", programName)
    }
		return true
}
