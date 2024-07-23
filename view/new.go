package view

import (
	"github.com/YumaFuu/ssm-tui/layout"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func CreateViewNew() *tview.Flex {
	// TypeとValueを入力できるようにする

	inputValue := tview.NewInputField().
		SetLabel("Value").
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorDefault)

	selectType := tview.NewDropDown().
		SetLabel("Type").
		SetFieldWidth(0).
		SetFieldBackgroundColor(tcell.ColorDefault).
		SetOptions([]string{"String", "SecureString", "StringList"}, nil)

	v := tview.NewFlex()

	v.
		AddItem(inputValue, 0, 1, false).
		AddItem(selectType, 0, 1, false).
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(true)

	return layout.ToModal(v, 80, 20)
}
