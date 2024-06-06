package main

import (
	"fmt"
	"log"

	beep "github.com/gen2brain/beeep"
)

// On application start, constantly checks specified folder for new files/folders,
// then parses them to check file size

type ApplicationData struct {
	name    string
	author  string
	license string
	version string
}

func Application() ApplicationData {
	return ApplicationData{
		name:    "VTApp",
		author:  "Breath",
		license: "Na",
		version: "0.0.1",
	}
}

func filePath() {
	// Find Downloads folder
}

func main() {
	msg := fmt.Sprintf("Application has started successfully...\n") // File path in msg
	err := beep.Alert(Application().name, msg, "path/to/icon")
	if beep.ErrUnsupported != nil {
		log.Fatal("Operating system not supported\n", err)
	}
}
