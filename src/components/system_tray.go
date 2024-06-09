package components

import (
	"log"
	"os"
	"path/filepath"

	"github.com/breathxv/vtapp/src/errors"
	widget "github.com/breathxv/vtapp/src/widget"
	"github.com/getlantern/systray"
)

// On ready defines the media, text and functionality
// that is attached to the system tray application icon.
func OnReady() {
	go DirectoryParse()
	bytes, err := os.ReadFile(filepath.Join("src", "assets", "VTApp.ico"))
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	systray.SetIcon(bytes)
	systray.SetTitle("VTApp")
	systray.SetTooltip("Intercepting downloads and scanning them...")
	changeDirectory := systray.AddMenuItem("Directory", "Change the directory the application detects from.")
	configWindow := systray.AddMenuItem("Config", "Open the configuration interface.")
	stopProgram := systray.AddMenuItem("Stop", "Stops the application but keeps it in the tray.")
	systray.AddSeparator()
	quitProgram := systray.AddMenuItemCheckbox("Quit", "Quit the application.", false)
	go func() {
		<-changeDirectory.ClickedCh

		log.Printf("User requested directory change...")
		w.Remove(set_directory)
		log.Printf("Stop watcher...")
		AdjustDirectory()
	}()
	go func() {
		<-configWindow.ClickedCh

		log.Printf("User requested config interface...")
		widget.ConfigInterface()
	}()
	go func() {
		<-stopProgram.ClickedCh
		if quitProgram.Checked() {
			for _, element := range w.WatchList() {
				err = w.Remove(element)
				if err != nil {
					log.Printf("Error removing path from event watcher.")
					errors.ToastAlert()
				}
			}
		} else {
			AdjustDirectory()
		}
	}()
	go func() {
		<-quitProgram.ClickedCh

		log.Printf("User requested application quit...")
		systray.Quit()
		os.Exit(0)
	}()
}

// On exit defines what's executed
// when the application is quitted.
func OnExit() {
	log.Fatalf("Quit from system tray.")
}
