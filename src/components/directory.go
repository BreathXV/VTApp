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
		log.Print("Watched directory set to path:", dir)
		DirWatcher(dir)
	} else {
		path := fmt.Sprintf("%s\\Downloads", dir)
		log.Print("Watched directory set to path:", path)
		DirWatcher(path)
	}
}
