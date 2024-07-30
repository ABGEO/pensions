package main

import (
	"os"

	"github.com/abgeo/pensions/internal/processor"
)

func main() {
	proc, err := processor.New()
	if err != nil {
		panic(err)
	}

	os.Exit(proc.Process())
}
