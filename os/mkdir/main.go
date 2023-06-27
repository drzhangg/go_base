package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.MkdirTemp(os.TempDir(), ""); err != nil {
		err = fmt.Errorf("failed to create a temp directory, error is %v", err)
		return
	}
}
