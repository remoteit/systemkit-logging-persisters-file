package persisters

import (
	"io"
	"os"

	logging "github.com/codemodify/systemkit-logging"
)

type fileLogger struct {
	file         *os.File
	errorOccured bool
	errorWriter  io.Writer
}

// NewFileLogger -
func NewFileLogger(fileName string, errorWriter io.Writer) logging.Logger {
	var f *os.File
	var err error

	if !fileOrFolderExists(fileName) {
		f, err = os.Create(fileName)
	} else {
		f, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0660)
	}
	if err != nil && errorWriter != nil {
		errorWriter.Write([]byte(err.Error() + "\n"))
	}

	return &fileLogger{
		file:         f,
		errorOccured: (err != nil),
		errorWriter:  errorWriter,
	}
}

func (thisRef fileLogger) Log(logEntry logging.LogEntry) logging.LogEntry {
	if thisRef.errorOccured && thisRef.errorWriter != nil {
		thisRef.errorWriter.Write([]byte(logEntry.Message + "\n"))
	} else {
		thisRef.file.WriteString(logEntry.Message + "\n")
	}

	return logEntry
}
