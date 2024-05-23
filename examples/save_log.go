package main

import (
	"os"

	"github.com/kaylee595/tracerr"
)

func main() {
	if err := read(); err != nil {
		// Save output to variable.
		text := tracerr.SprintSource(err)
		os.WriteFile("/tmp/tracerr.log", []byte(text), 0644)
	}
}

func read() error {
	return readNonExistent()
}

func readNonExistent() error {
	_, err := os.ReadFile("/tmp/non_existent_file")
	return tracerr.Wrap(err)
}
