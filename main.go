package main

import (
	"os"

	"github.com/xpo-nential/share-golang/logging"
)

func init() {
	// setup logger
	os.Setenv("LOG_LEVEL", "debug")
	logging.Init("share-golang")
}

func main() {
	logging.Info("Hello, World!")
	logging.Warn("Hello, World!")
	logging.Error("Hello, World!")
	logging.Debug("Hello, World!")
	logging.Fatal("Hello, World!")
}
