package main

import (
	cmp "github.com/breathxv/vtapp/src/components"
	"github.com/getlantern/systray"
)

// On application start, constantly checks specified folder for new files/folders,
// then parses them to check file size
// TODO: Finish this ^
func main() {
	// TODO: Welcome screen on initial start
	// Start system tray
	systray.Run(cmp.OnReady, cmp.OnExit)
}
