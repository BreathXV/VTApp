package handlers

// Sends files < 30MB to VirusTotal API

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/VirusTotal/vt-go"
	"github.com/breathxv/vtapp/src/errors"
	"github.com/gen2brain/beeep"
)

// TODO: Add docs for var here!
var (
	apikey = flag.String(
		"apikey",
		os.Getenv("VIRUSTOTAL_API_KEY"),
		"VirusTotal API Key",
	)
)

// Parse files takes a file or folder and hands it over to the Virus Total API
// for scanning. It takes one argument, path (string) and returns err (error).
// Path should be the location of the folder/file that is to be scanned.
// The folder cannot be a directory, it must be zipped before passing to the
// VirusTotal API. Error is only returned in the event the function has an error.
// The API will send an object back that is handed to the file report, which extracts
// values from it and prompts them to the user through Toast.
func ParseFile(path string) (err error) {
	flag.Parse()

	if *apikey == "" {
		// Text box to set API Key
	}

	client := vt.NewClient(*apikey)

	filePointer, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		errors.ToastAlert()
	}

	file, err := os.Stat(path)
	if err != nil {
		errors.ToastAlert()
	}
	fmt.Print(file.Name())

	results, err := client.NewFileScanner().ScanFile(
		filePointer,
		nil,
	)
	if err != nil {
		beeep.Alert(
			"VTApp",
			"Failed to send files to be scanned!",
			filepath.Join("src", "assets", "VTApp.ico"),
		)
	}

	fmt.Print(results.Attributes())
	return nil
}
