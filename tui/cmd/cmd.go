package cmd

import (
	"github.com/YumaFuu/s1m/tui/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type (
	CmdBox struct {
		*tview.InputField
		pubsub *pubsub.PubSub
	}

	ConfirmInput struct {
		Label     string
		Successor func()
	}
)

func NewCmdBox(ps *pubsub.PubSub) *CmdBox {
	style := tcell.StyleDefault.Background(tcell.ColorReset)
	style = style.Attributes(tcell.AttrBold)

	v := tview.NewInputField().
		SetFieldBackgroundColor(tcell.ColorDefault).
		// TODO: ColorDefault is not working
		SetLabelStyle(style).
		SetPlaceholderStyle(tcell.StyleDefault.Foreground(tcell.ColorGray))

	v.
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(false)

	return &CmdBox{v, ps}
}

func (v *CmdBox) Confirm(ci ConfirmInput) {
	v.SetLabel(ci.Label + " (y/n): ")
	v.SetPlaceholder("")
	v.pubsub.Pub(nil, pubsub.TopicAppDraw)

	v.SetDoneFunc(func(key tcell.Key) {
		t := v.GetText()
		if key == tcell.KeyEnter {
			if t == "y" || t == "Y" || t == ":w" {
				ci.Successor()
			}
		}

		v.SetLabel("")
		v.SetText("")
		v.pubsub.Pub(nil, pubsub.TopicAppReload)
	})
}
