package vbox

import (
	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	ModeUpdate Mode = iota
	ModeCreate
	ModeDefault = ModeUpdate
)

const (
	KeyInput    = "KeyInput"
	KeyTextArea = "KeyTextArea"
	KeyList     = "KeyList"
)

type (
	Mode     int
	ValueBox struct {
		*tview.Pages
		TextArea *tview.TextArea
		Input    *tview.InputField
		ListView *tview.List
		pubsub   *pubsub.PubSub
		param    ssm.Parameter
		mode     Mode
	}
)

func NewValueBox(ps *pubsub.PubSub) *ValueBox {
	textArea := tview.
		NewTextArea().
		SetTextStyle(
			tcell.
				StyleDefault.
				Foreground(tcell.ColorWhite).
				Background(tcell.ColorDefault),
		)
	textArea.
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(true)

	input := tview.NewInputField()
	s := tcell.StyleDefault.Background(tcell.ColorReset)
	input.
		SetLabelStyle(s).
		SetFieldStyle(s).
		SetBorder(true).
		SetBackgroundColor(tcell.ColorDefault)

	listView := tview.NewList().
		SetMainTextStyle(s).
		SetSecondaryTextStyle(s).
		SetShortcutStyle(s).
		SetSelectedStyle(
			s.
				Foreground(tcell.ColorWhite).
				Bold(true).
				Underline(true),
		)

	listView.
		SetBorder(true).
		SetBackgroundColor(tcell.ColorDefault)

	listView.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'j':
			listView.SetCurrentItem(listView.GetCurrentItem() + 1)
		case 'k':
			listView.SetCurrentItem(listView.GetCurrentItem() - 1)
		}
		return event
	})

	pages := tview.NewPages().
		AddPage(KeyTextArea, textArea, true, true).
		AddPage(KeyInput, input, true, false).
		AddPage(KeyList, listView, true, false)

	vbox := &ValueBox{
		pages,
		textArea,
		input,
		listView,
		ps,
		ssm.Parameter{},
		ModeDefault,
	}

	return vbox
}

func (v *ValueBox) SetPrev(s ssm.Parameter) {
	v.param = s
}

func (v *ValueBox) GetPrev() ssm.Parameter {
	return v.param
}

func (v *ValueBox) SetMode(m Mode) {
	v.mode = m
}

func (v *ValueBox) SetParam(p ssm.Parameter) {
	v.param = p
}

func (v *ValueBox) GetText() string {
	return v.TextArea.GetText()
}

func (v *ValueBox) SetText(s string, b bool) {
	v.TextArea.SetText(s, b)
}

func (v *ValueBox) WaitTopic() {
	chUpdate := v.pubsub.Sub(pubsub.TopicWriteValueBox)

	for msg := range chUpdate {
		v.SetMode(ModeUpdate)

		if p, ok := msg.(ssm.Parameter); ok {
			s := *p.Value

			v.TextArea.SetText(s, true)
			v.SetPrev(p)
		}
		if b, ok := msg.(string); ok {
			v.TextArea.SetText(b, true)
		}
	}
}
