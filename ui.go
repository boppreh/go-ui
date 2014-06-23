package main

import (
	"github.com/boppreh/forked-go-ui"
	"time"
)

func Progress(title string, descriptionChannel chan string, progressChannel chan int) {
	w := ui.NewWindow(title, 400, 120)

	descriptionLabel := ui.NewLabel("")
	progressBar := ui.NewProgressBar()
	w.Open(ui.Layout(descriptionLabel, progressBar))

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
		}
	}
}

func myMain() {
	descriptionChannel := make(chan string)
	progressChannel := make(chan int)
	go func() {
		for i := 1; i <= 100; i++ {
			progressChannel <- i
			descriptionChannel <- "Doing " + string(i)
			time.Sleep(time.Millisecond * 100)
		}
		close(progressChannel)
	}()
	Progress("Title", descriptionChannel, progressChannel)
}
