package main

import (
	"os"
	"path/filepath" 
	"log"
	"fmt"
)

func sendingToLogfile(stringToLog string) {
	// Gathering the file location of the logfile
	var mainFolderDir string = findSyncNestFolderDir()
	var logfileLocation string = filepath.Join(mainFolderDir, "logfile.log")

	// Open the logfile in append mode --> Create it if it DNE
	logFile, _ := os.OpenFile(logfileLocation, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) // Can't check for error --> No way to log
	defer logFile.Close()

	// Making the default log output the logfile
	log.SetOutput(logFile)

	// Sending String to Logfile --> NOTE: log module automatically generates dt boilerplate entry
	stringToLog = fmt.Sprintf("\n%s", stringToLog)

	// Logging to the logfile.log
	log.Println(stringToLog)
}