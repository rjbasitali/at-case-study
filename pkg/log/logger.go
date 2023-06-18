package log

import (
	"os"

	"github.com/rjbasitali/go-log"
)

var logger log.Logger

func init() {
	logger = log.NewLogger(os.Stdout, os.Stderr)
}

func Log(args ...interface{}) {
	logger.Log(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}
