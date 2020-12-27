package main

import (
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func startSystray() {
	defer func() {
		recover()
	}()
	systray.SetIcon(icon.Data)
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()

}

func onExit() {
	// clean up here
	os.Exit(0)
}
