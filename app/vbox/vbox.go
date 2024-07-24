package vbox

import (
	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ValueBox struct {
	*tview.TextArea
	pubsub *pubsub.PubSub
	prev   string
}

func NewValueBox(ps *pubsub.PubSub) *ValueBox {
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

	return &ValueBox{v, ps, ""}
}

func (v *ValueBox) SetPrev(s string) {
	v.prev = s
}

func (v *ValueBox) GetPrev() string {
	return v.prev
}

func (v *ValueBox) WaitTopic() {
	chUpdate := v.pubsub.Sub(pubsub.TopicUpdateValueBox)
	chUpdateBorder := v.pubsub.Sub(pubsub.TopicUpdateValueBoxBorder)

	for {
		select {
		case msg := <-chUpdate:
			if s, ok := msg.(string); ok {
				v.SetText(s, true)
				v.SetPrev(s)
			}
		case msg := <-chUpdateBorder:
			if b, ok := msg.(tcell.Color); ok {
				v.SetBorderColor(b)
			}
		}
	}
}
