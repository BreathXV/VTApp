package widget

import (
	"log"
	"os"
	"path/filepath"

	"github.com/breathxv/vtapp/src/errors"
	ui "github.com/ncruces/zenity"
)

func DirWindow() (path string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		errors.ToastAlert()
		log.Fatal("Cannot find home directory.")
	}

	path, err = ui.SelectFile(
		ui.Filename(homeDir),
		ui.WindowIcon(filepath.Join("src", "assets", "VTApp.ico")),
		ui.Directory(),
	)
	if err != nil {
		log.Fatal(err)
	}

	return path
}
