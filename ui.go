package goui

import (
	"github.com/boppreh/forked-go-ui"
)

func Progress(title string, descriptionChannel chan string, progressChannel chan int, cancel func()) {
	w := ui.NewWindow(title, 383, 143)

	descriptionLabel := ui.NewLabel("")
	progressBar := ui.NewProgressBar()
	cancelButton := ui.NewButton("Cancel")
	layout := ui.Layout(descriptionLabel, progressBar, ui.Layout(nil, cancelButton))
	w.Open(layout)

	for {
		select {
		case d, ok := <-descriptionChannel:
			if !ok {
				return
			}
			descriptionLabel.SetText(d)
		case p, ok := <-progressChannel:
			if !ok {
				return
			}
			progressBar.SetProgress(p)
		case <-w.Closing:
			return
		case <- cancelButton.Clicked:
			if cancel == nil {
				return
			} else {
				go cancel()
			}
		}
	}
}

func Error(primaryText string, secondaryText string) {
	ui.MsgBoxError(primaryText, secondaryText)
}

func Info(primaryText string, secondaryText string) {
	ui.MsgBox(primaryText, secondaryText)
}

func Start(main func()) {
	err := ui.Go(main)
	if err != nil {
		panic(err)
	}
}
