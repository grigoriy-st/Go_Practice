package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("First Window")

	hello := widget.NewLabel("My first app")

	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Start", func() {
			hello.SetText("Aaa, welcome! :)")
		}),
	))
	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyQ, Modifier: fyne.KeyModifierControl}

	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		w.Close()
		log.Println("App closed")
	})

	w.ShowAndRun()
}
