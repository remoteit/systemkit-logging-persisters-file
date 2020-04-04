package persisters

import (
	"fmt"
	"os"

	logging "github.com/codemodify/systemkit-logging"
)

func fileOrFolderExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

type emptyWritter struct{}

func (thisRef emptyWritter) Write(p []byte) (n int, err error) {
	return 0, nil
}

// NewFileLoggerCustomName -
func NewFileLoggerCustomName(fileName string) logging.Logger {
	return NewFileLogger(fileName, &emptyWritter{})
}

// NewFileLoggerDefaultName -
func NewFileLoggerDefaultName() logging.Logger {
	return NewFileLoggerCustomName(fmt.Sprintf("%s.log", os.Args[0]))
}

// NewFileLoggerCustomNameEasy -
func NewFileLoggerCustomNameEasy(fileName string) logging.LoggerImplementation {
	return logging.NewDefaultLoggerImplementation(NewFileLoggerCustomName(fileName))
}

// NewFileLoggerDefaultNameEasy -
func NewFileLoggerDefaultNameEasy() logging.LoggerImplementation {
	return NewFileLoggerCustomNameEasy(fmt.Sprintf("%s.log", os.Args[0]))
}
