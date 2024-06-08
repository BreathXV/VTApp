package widget

import (
	"log"
	"os"

	"github.com/breathxv/vtapp/src/errors"
	ui "github.com/ncruces/zenity"
)

// ! Doesn't work in system tray
func DirWindow() (path string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		errors.ToastAlert()
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
