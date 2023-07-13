package helpers

import (
	"fmt"
	"log"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// https://github.com/gin-gonic/gin/issues/1376
func init() {
	logfile, err := os.Create("log/sala.log")
	if err != nil {
		fmt.Println("Open Log Sala Failed", err)
	}

	Info = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(logfile, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logfile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
