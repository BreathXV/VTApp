package widget

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	A fyne.App
	W fyne.Window
)

// TODO: Create UI for profile creation
func CreateProfile() {
	A = app.New()
	log.Printf("Created new Fyne app.")
	W = A.NewWindow("VTApp - Profile Builder")
	log.Printf("Created window.")

	profile := widget.NewLabel("This is the profile builder!")
	log.Printf("Added label.")
	W.SetContent(container.NewVBox(
		profile,
		widget.NewButton("Hi!", func() {
			profile.SetText("Welcome :)!")
		}),
	))
	log.Printf("Set content.")
	W.Show()
}
