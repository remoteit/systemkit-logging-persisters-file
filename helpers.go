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

// NewFileLoggerDefaultName -
func NewFileLoggerDefaultName() logging.Logger {
	return NewFileLogger(fmt.Sprintf("%s.log", os.Args[0]), &emptyWritter{})
}

type emptyWritter struct{}

func (thisRef emptyWritter) Write(p []byte) (n int, err error) {
	return 0, nil
}
