package main

import (
	cmp "github.com/breathxv/vtapp/src/components"
	"github.com/getlantern/systray"
)

// On application start, constantly checks specified folder for new files/folders,
// then parses them to check file size

func main() {
	// TODO: Welcome screen on initial start
	systray.Run(cmp.OnReady, cmp.OnExit)
	// Start directory selection
	cmp.DirectoryParse()
}
