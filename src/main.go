package main

import (
	"fmt"
	"log"
	"os"

	cmp "github.com/breathxv/vtapp/src/components"
	ui "github.com/ncruces/zenity"
)

// On application start, constantly checks specified folder for new files/folders,
// then parses them to check file size

func dirWindow() (path string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		cmp.ToastAlert()
		log.Fatal("Cannot find home directory.")
	}

	path, err = ui.SelectFile(
		ui.Filename(homeDir),
		ui.WindowIcon(`.\assets\RW.png`),
		ui.Directory(),
	)
	if err != nil {
		log.Fatal(err)
	}

	return path
}

func main() {
	dir, err := os.UserHomeDir()
	if err != nil {
		dir = dirWindow()
		cmp.DirWatcher(dir)
	} else {
		path := fmt.Sprintf("%s\\Downloads", dir)
		cmp.DirWatcher(path)
	}
}
