package components

import (
	"fmt"
	"log"
	"os"

	ui "github.com/ncruces/zenity"
)

func DirWindow() (path string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		ToastAlert()
		log.Fatal("Cannot find home directory.")
	}

	path, err = ui.SelectFile(
		ui.Filename(homeDir),
		ui.WindowIcon(`.\assets\VTApp.ico`),
		ui.Directory(),
	)
	if err != nil {
		log.Fatal(err)
	}

	return path
}

func DirectoryParse() {
	dir, err := os.UserHomeDir()
	if err != nil {
		dir = DirWindow()
		DirWatcher(dir)
	} else {
		path := fmt.Sprintf("%s\\Downloads", dir)
		DirWatcher(path)
	}
}
