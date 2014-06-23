package main

import (
	"github.com/boppreh/forked-go-ui"
)

func myMain() {
	w := ui.NewWindow("Test Window", 400, 250)
	ui.AppQuit = w.Closing
	w.Open(ui.Layout(ui.NewLabel("Task Description"),
		ui.NewProgressBar(),
		nil,
		ui.Layout(ui.NewButton("Button 1"),
			ui.NewLineEdit("Edit"),
			nil,
			ui.Layout(ui.NewButton("Button 2"),
				ui.NewButton("Button 3"),
				ui.NewButton("Button 4")),
			ui.Layout(ui.NewButton("Button 5"),
				nil,
				ui.NewButton("Button 6")))))

	for {
		select {
		case <-w.Closing:
			return
		}
	}
}

func main() {
	err := ui.Go(myMain)
	if err != nil {
		panic(err)
	}
}
