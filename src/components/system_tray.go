package components

import (
	"log"
	"os"

	"github.com/getlantern/systray"
)

func OnReady() {
	bytes, _ := os.ReadFile("src\\assets\\VTApp.ico")
	systray.SetTemplateIcon(bytes, bytes)
	systray.SetTitle("VTApp")
	quitProgram := systray.AddMenuItem("Quit", "Quit the application.")
	changeDirectory := systray.AddMenuItem(
		"Directory",
		"Change the directory the application detects from.",
	)
	go func() {
		<-changeDirectory.ClickedCh
		log.Printf("User requested directory change...")
		w.Close()
		log.Printf("Stop watcher...")
		DirectoryParse()
	}()
	go func() {
		<-quitProgram.ClickedCh
		log.Printf("User requested application quit...")
		systray.Quit()
		os.Exit(0)

	}()
}

func OnExit() {
	log.Fatalf("Quit from system tray.")
}
