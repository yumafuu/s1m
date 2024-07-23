package vbox

import (
	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ValueBox struct {
	*tview.TextArea
	ch chan string
}

func NewValueBox(ps pubsub.PubSub) ValueBox {
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

	ch := ps.Sub(pubsub.TopicUpdateValueBox)

	return ValueBox{v, ch}
}

func (v ValueBox) WaitTopic() {
	for msg := range v.ch {
		v.SetText(msg, true)
	}
}
