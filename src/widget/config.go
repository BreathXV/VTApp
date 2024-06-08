package widget

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func ConfigInterface() {
	a := app.New()
	w := a.NewWindow("VTApp")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
