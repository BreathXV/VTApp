package main

import (
	"fmt"
	"log"
	"os"

	ui "github.com/ncruces/zenity"
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
	dir, err := os.UserHomeDir()
	if err != nil {
		dir = ``

	}

	path, err := ui.SelectFile(
		ui.Filename(fmt.Sprintf(`%s\Downloads`, dir)),
		ui.WindowIcon(`.\assets\RW.png`),
		ui.Directory(),
	)
	if err != nil {
		log.Fatal(err)
	}
}
