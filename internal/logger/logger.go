package logger

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func Logger() (*os.File, error) {
	err := os.MkdirAll("logs", 0o755)
	if err != nil {
		return nil, err
	}

	Log_file := "logs/server.log"
	file, err := os.OpenFile(Log_file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return nil, err
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return file, nil
}
