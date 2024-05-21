// Package osutil provides reusable functions not available in the standard library's os package.
package osutil

import (
	"os"
	"os/exec"
	"path/filepath"
)

// CheckIfFileExists checks if a file exists on the PATH, or at a given absolute path.
func CheckIfProgramExists(program string) bool {
	// Check if program path is absolute
	if filepath.IsAbs(program) {
		_, err := os.Stat(program)
		return !os.IsNotExist(err)
	}

	// Check if program exists in PATH
	_, err := exec.LookPath(program)
	return err == nil
}
