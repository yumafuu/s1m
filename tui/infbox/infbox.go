package infbox

import (
	"fmt"

	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	ValueFormat = `Version:          %d
Name:             %s
Type:             %s
LastModifiedDate: %s`
	UpdateMessageFormat = `[blue]New Parameter Updated
Name:             %s
Type:             %s
Value:            %s`
	CreateMessageFormat = `[green]New Parameter Created
Name:             %s
Type:             %s
Value:            %s`
	DeleteMessageFormat = `[red]New Parameter Deleted
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
	chWrite := v.pubsub.Sub(pubsub.TopicWriteInfoBox)

	for {
		msg := <-chWrite
		if s, ok := msg.(string); ok {
			v.SetText(s)
		} else if s, ok := msg.(ssm.Parameter); ok {
			var name string
			if s.Name == nil {
				name = ""
			} else {
				name = *s.Name
			}

			var t string
			if s.LastModifiedDate == nil {
				t = ""
			} else {
				t = s.LastModifiedDate.String()
			}

			fs := fmt.Sprintf(
				ValueFormat,
				s.Version,
				name,
				s.Type,
				t,
			)
			v.SetText(fs)
		}
	}
}
