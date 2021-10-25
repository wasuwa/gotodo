package logging

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("file=%s err=%s", logFile, err.Error())
	}
	multLogFile := io.MultiWriter(os.Stdout, f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multLogFile)
}
