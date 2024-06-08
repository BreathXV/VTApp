package errors

import (
	"log"

	toast "github.com/gen2brain/beeep"
)

// Toast alert is called when a function within
// the application failed which is crucial to its
// functionality (i.e. the program cannot function
// without it). It sends a Windows Toast notification
// to the user that displays an error message and then,
// quits the program.
func ToastAlert() {
	tErr := toast.Alert(
		"VTApp",
		"An error has occurred.\nPlease refer to the logs for further information.",
		"src\\assets\\VTApp.ico",
	)
	if tErr != nil {
		log.Fatal("Failed to dispatch Toast.")
	}
}
