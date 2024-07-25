package vbox

import (
	"github.com/YumaFuu/ssm-tui/tui/pubsub"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	ModeUpdate Mode = iota
	ModeCreate
)

const ModeDefault = ModeUpdate

type (
	Mode     int
	ValueBox struct {
		*tview.TextArea
		pubsub *pubsub.PubSub
		param  types.Parameter
		mode   Mode
	}
)

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

	return &ValueBox{v, ps, types.Parameter{}, ModeDefault}
}

func (v *ValueBox) SetPrev(s types.Parameter) {
	v.param = s
}

func (v *ValueBox) GetPrev() types.Parameter {
	return v.param
}

func (v *ValueBox) SetMode(m Mode) {
	v.mode = m
}

func (v *ValueBox) SetParam(p types.Parameter) {
	v.param = p
}

func (v *ValueBox) WaitTopic() {
	chUpdate := v.pubsub.Sub(pubsub.TopicUpdateValueBox)
	chUpdateBorder := v.pubsub.Sub(pubsub.TopicUpdateValueBoxBorder)

	for {
		select {
		case msg := <-chUpdate:
			v.SetMode(ModeUpdate)

			if p, ok := msg.(types.Parameter); ok {
				s := *p.Value
				v.SetText(s, true)
				v.SetPrev(p)
			}
			if b, ok := msg.(string); ok {
				v.SetText(b, true)
			}
		case msg := <-chUpdateBorder:
			if b, ok := msg.(tcell.Color); ok {
				v.SetBorderColor(b)
			}
		}
	}
}
