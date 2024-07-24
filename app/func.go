package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *App) SetInputCapture() *tview.Application {
	inputCapture := func(event *tcell.EventKey) *tcell.EventKey {
		switch a.tapp.GetFocus() {
		case a.ptree:
			a.ptree.InputCapture(event)
		case a.vbox:
			a.vbox.InputCapture(event)
		}

		if event.Key() == tcell.KeyCtrlC {
			a.tapp.Stop()
		}

		return event
	}
	a.tapp.SetInputCapture(inputCapture)

	return a.tapp
}
