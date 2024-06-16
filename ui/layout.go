package ui

import (
	"github.com/rivo/tview"
)

func BuildLayout(
	parameterTree *tview.TreeView,
	infoView *tview.TextView,
	valueView *tview.TextView,
) *tview.Flex {

	param := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(infoView, 0, 1, false).
		AddItem(valueView, 0, 4, false)

	flex := tview.
		NewFlex().
		AddItem(parameterTree, 0, 1, true).
		AddItem(param, 0, 2, false)

	layout := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(flex, 0, 1, true)

	return layout
}
