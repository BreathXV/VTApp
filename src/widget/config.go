package widget

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

// ? May be removed/fully re-worked
func ConfigInterface() {
	a := app.New()
	w := a.NewWindow("VTApp")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
