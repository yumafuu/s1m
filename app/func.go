package app

import (
	"github.com/gdamore/tcell/v2"
)

func (a *App) InputCapture(event *tcell.EventKey) *tcell.EventKey {

	switch a.GetFocus() {
	case a.ptree:
		a.ptree.InputCapture(event)
	case a.vbox:
		a.vbox.InputCapture(event)
	}

	if event.Key() == tcell.KeyCtrlC {
		a.Stop()
	}

	return event
}
