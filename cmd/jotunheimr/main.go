package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"jotunheimr/internal/xconnector"
)

func main() {
	a := app.New()
	w := a.NewWindow("jotunheimr")

	conn := xconnector.NewConn()

	w.SetContent(container.NewVBox(
		widget.NewButton("Connect", conn.Connect),
	))

	w.ShowAndRun()
}
