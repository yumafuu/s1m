package infbox

import (
	"github.com/YumaFuu/s1m/tui/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	ValueFormat = `Version:          %d
Name:             %s
Type:             %s
LastModifiedDate: %s`
	CreateMessageFormat = `[green]New Parameter Created
Name:             %s
Type:             %s
Value:            %s`
	UpdateMessageFormat = `[green]New Parameter Updated
Name:             %s
Type:             %s
Value:            %s`
	DeleteMessageFormat = `[green]New Parameter Deleted
Name:             %s`
)

type InfoBox struct {
	*tview.TextView
	pubsub *pubsub.PubSub
}

func NewInfoBox(ps *pubsub.PubSub) *InfoBox {
	v := tview.NewTextView().
		SetDynamicColors(true).
		SetWrap(true)

	v.
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(true)

	return &InfoBox{v, ps}
}

func (v InfoBox) WaitTopic() {
	chUpdate := v.pubsub.Sub(pubsub.TopicUpdateInfoBox)

	for {
		msg := <-chUpdate
		if s, ok := msg.(string); ok {
			v.SetText(s)
		}
	}
}
