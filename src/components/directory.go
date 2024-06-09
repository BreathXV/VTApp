package components

import (
	"log"
	"os"
	"path/filepath"

	widget "github.com/breathxv/vtapp/src/widget"
)

var (
	set_directory string
)

// TODO: Add a check that the program defined path is correct.
// Directory Parse attempts to assign the users download path.
// It does this by first checking if their home directory exists,
// if it does, then it will automatically attach "Downloads" onto
// the end of the path. Otherwise, if the OS cannot locate the
// home directory, then it call to DirWindow() which will return
// the user's selected path. Once the path has been defined, it
// will hand it to the DirWatcher().
func DirectoryParse() {
	dir, err := os.UserHomeDir()
	if err != nil {
		dir = widget.DirWindow()
		log.Print("Watched directory set to path:", dir)
		set_directory = dir
		DirWatcher(dir)
	} else {
		path := filepath.Join(dir, "Downloads")
		log.Print("Watched directory set to path:", path)
		set_directory = path
		DirWatcher(path)
	}
}
