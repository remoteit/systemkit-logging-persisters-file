package persisters

import (
	"io"
	"os"

	logging "github.com/codemodify/systemkit-logging"
)

type fileLogger struct {
	file          *os.File
	errorOccurred bool
	errorWriter   io.Writer
}

// NewFileLogger -
func NewFileLogger(fileName string, errorWriter io.Writer) logging.CoreLogger {
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
		file:          f,
		errorOccurred: (err != nil),
		errorWriter:   errorWriter,
	}
}

func (thisRef fileLogger) Log(logEntry logging.LogEntry) logging.LogEntry {
	if thisRef.errorOccurred && thisRef.errorWriter != nil {
		thisRef.errorWriter.Write([]byte(logEntry.Message + "\n"))
	} else {
		thisRef.file.WriteString(logEntry.Message + "\n")
	}

	return logEntry
}
