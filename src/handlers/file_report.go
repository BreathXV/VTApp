package handlers

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	data "github.com/breathxv/vtapp/src/data"
	"github.com/gen2brain/beeep"
)

func GetFileReport(file_hash string) (fileReport data.FileReportStructure) {
	url := "https://www.virustotal.com/api/v3/files/id"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		beeep.Alert(
			"VTApp",
			"Failed to create request to VirusTotal.",
			filepath.Join("src", "assets", "VTApp.ico"),
		)
	}

	req.Header.Add("accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		beeep.Alert(
			"VTApp",
			"Failed to create request to VirusTotal.",
			filepath.Join("src", "assets", "VTApp.ico"),
		)
	}

	defer res.Body.Close()
	// TODO: Body as JSON
	// fileReport, err = io.ReadAll(res.Body)

	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&fileReport)
	if err != nil {
		beeep.Alert(
			"VTApp",
			"Failed to decode received file report from Virus Total.",
			filepath.Join("src", "assets", "VTApp.ico"),
		)
	}
	return fileReport
}
