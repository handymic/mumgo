package mumgo

import (
	"fmt"
	"os"
	"path/filepath"
)

// Derives the current working directory of running program
func pwd() string {
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		fmt.Errorf("cannot get working directory: %s", err)
		os.Exit(1)
	}

	return pwd
}
