package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildValueView() *tview.TextArea {
	v := tview.
		NewTextArea().
		SetTextStyle(
			tcell.
				StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault),
		)
	v.
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(true)

	return v
}
