package infbox

import (
	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type InfoBox struct {
	*tview.TextView
	ch chan string
}

func NewInfoBox(ps pubsub.PubSub) InfoBox {
	v := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true)

	v.
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(true)

	ch := ps.Sub(pubsub.TopicUpdateInfoBox)

	return InfoBox{v, ch}
}

func (v InfoBox) WaitTopic() {
	for msg := range v.ch {
		v.SetText(msg)
	}
}
