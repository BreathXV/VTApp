package components

import (
	"fmt"
	"log"
	"os"

	widget "github.com/breathxv/vtapp/src/widget"
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
		DirWatcher(dir)
	} else {
		path := fmt.Sprintf("%s\\Downloads", dir)
		log.Print("Watched directory set to path:", path)
		DirWatcher(path)
	}
}
