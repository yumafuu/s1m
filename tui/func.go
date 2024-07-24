package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *Tui) SetInputCapture() *tview.Application {
	inputCapture := func(event *tcell.EventKey) *tcell.EventKey {
		switch a.app.GetFocus() {
		case a.ptree:
			a.ptree.InputCapture(event)
		case a.vbox:
			a.vbox.InputCapture(event)
		}

		if event.Key() == tcell.KeyCtrlC {
			a.app.Stop()
		}

		return event
	}
	a.app.SetInputCapture(inputCapture)

	return a.app
}
