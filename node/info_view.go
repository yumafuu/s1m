package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildInfoView() *tview.TextView {
	// Create a text view to display the selected node's value
	v := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true)

	v.
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(true)

	return v
}
