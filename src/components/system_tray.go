package components

import (
	"log"
	"os"

	"github.com/getlantern/systray"
)

func OnReady() {
	bytes, err := os.ReadFile("src\\assets\\VTApp.ico")
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	systray.SetIcon(bytes)
	log.Print("Set icon.")
	systray.SetTitle("VTApp")
	log.Print("Set title.")
	quitProgram := systray.AddMenuItem("Quit", "Quit the application.")
	log.Print("Add menu item: Quit")
	// ! Fix error on attempted directory change
	// changeDirectory := systray.AddMenuItem(
	// 	"Directory",
	// 	"Change the directory the application detects from.",
	// )
	// log.Print("Added menu item: Directory")
	// go func() {
	// 	<-changeDirectory.ClickedCh
	// 	log.Printf("User requested directory change...")
	// 	log.Printf("Stop watcher...")
	// 	DirectoryParse()
	// }()
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
