package persisters

import (
	"fmt"
	"os"

	logging "github.com/remoteit/systemkit-logging"
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
func NewFileLoggerCustomName(fileName string) logging.CoreLogger {
	return NewFileLogger(fileName, &emptyWritter{})
}

// NewFileLoggerDefaultName -
func NewFileLoggerDefaultName() logging.CoreLogger {
	return NewFileLoggerCustomName(fmt.Sprintf("%s.log", os.Args[0]))
}

// NewFileLoggerCustomNameEasy -
func NewFileLoggerCustomNameEasy(fileName string) logging.Logger {
	return logging.NewLoggerImplementation(NewFileLoggerCustomName(fileName))
}

// NewFileLoggerDefaultNameEasy -
func NewFileLoggerDefaultNameEasy() logging.Logger {
	return NewFileLoggerCustomNameEasy(fmt.Sprintf("%s.log", os.Args[0]))
}
