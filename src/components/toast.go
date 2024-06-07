package components

import (
	"log"

	toast "github.com/gen2brain/beeep"
)

func ToastAlert() {
	tErr := toast.Alert(
		"VTApp",
		"An error has occurred.\nPlease refer to the logs for further information.",
		"src\\assets\\RW.ico",
	)
	if tErr != nil {
		log.Fatal("Failed to dispatch Toast.")
	}
}
