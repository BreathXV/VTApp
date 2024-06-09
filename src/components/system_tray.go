package components

import (
	"log"
	"os"
	"path/filepath"

	widget "github.com/breathxv/vtapp/src/widget"
	"github.com/getlantern/systray"
)

// On ready defines the media, text and
// functionality that is attached to the
// system tray application icon.
func OnReady() {
	go DirectoryParse()
	bytes, err := os.ReadFile(filepath.Join("src", "assets", "VTApp.ico"))
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	systray.SetIcon(bytes)
	systray.SetTitle("VTApp")
	configWindow := systray.AddMenuItem("Config", "Open the configuration interface.")
	quitProgram := systray.AddMenuItem("Quit", "Quit the application.")
	// ! Error on attempted directory change
	// ! Possible cause is attempting to pull
	// ! 'w' when it is in another thread.
	// TODO: Fix error
	changeDirectory := systray.AddMenuItem(
		"Directory",
		"Change the directory the application detects from.",
	)
	log.Print("Added menu item: Directory")
	go func() {
		<-changeDirectory.ClickedCh
		log.Printf("User requested directory change...")
		w.Remove(set_directory)
		log.Printf("Stop watcher...")
		DirectoryParse()
	}()
	go func() {
		<-configWindow.ClickedCh
		log.Printf("User requested config interface...")
		widget.ConfigInterface()
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
