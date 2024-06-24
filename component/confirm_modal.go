package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildConfirmModalView() *tview.Flex {
	modal := func(p tview.Primitive, width, height int) *tview.Flex {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	box := tview.NewBox().
		SetBorder(true).
		SetBackgroundColor(tcell.ColorDefault).
		SetBorderAttributes(tcell.AttrBold).
		SetBorderColor(tcell.ColorBlue)

	v := modal(box, 80, 20)

	return v
}
