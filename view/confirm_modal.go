package view

import (
	"github.com/YumaFuu/ssm-tui/layout"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildConfirmModalView() *tview.Flex {
	box := tview.NewBox().
		SetBorder(true).
		SetBackgroundColor(tcell.ColorDefault).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorBlue)

	v := layout.ToModal(box, 80, 20)

	return v
}
