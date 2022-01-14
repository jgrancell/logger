package main

import (
	"fmt"
	"os"

	"github.com/jgrancell/logger"
)

// You can use the `LogAndExit()` function to cleanly exit your app with error code 0 or 1
func main() {
	log := &logger.Logger{}
	_ = log.Init()

	os.Exit(log.LogAndExit(Run(log)))
}

func Run(log *logger.Logger) error {
	// This is where we will perform our actual app code
	err := fmt.Errorf("simulating an example error here")
	if err != nil {
		return err
	}
	return nil
}
