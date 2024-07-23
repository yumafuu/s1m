package app

import (
	"github.com/YumaFuu/ssm-tui/app/infbox"
	"github.com/YumaFuu/ssm-tui/app/ptree"
	"github.com/YumaFuu/ssm-tui/app/vbox"
	"github.com/rivo/tview"
)

type Layout struct {
	*tview.Pages
}

func NewLayout(
	parameterTree ptree.ParameterTree,
	infoBox infbox.InfoBox,
	valueBox vbox.ValueBox,
) Layout {
	param := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(infoBox, 0, 1, false).
		AddItem(valueBox, 0, 4, false)

	flex := tview.
		NewFlex().
		AddItem(parameterTree, 0, 1, true).
		AddItem(param, 0, 2, false)

	layout := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(flex, 0, 1, true)

	pages := tview.NewPages().
		AddPage("main", layout, true, true)
		// AddPage("new", createView, true, false)

	return Layout{pages}
}
