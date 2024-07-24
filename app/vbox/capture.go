package vbox

import (
	"fmt"

	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/gdamore/tcell/v2"
)

func (vbox *ValueBox) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		vbox.pubsub.Pub(true, pubsub.TopicSetAppFocusTree)
		vbox.pubsub.Pub("ESC is pressed", pubsub.TopicAppDraw)
		vbox.SetBorderColor(tcell.ColorDefault)

		prev := vbox.GetPrev()
		prevValue := *prev.Value
		newValue := vbox.GetText()
		if prevValue != newValue {
			p := ssm.Parameter{
				Name:  prev.Name,
				Value: &newValue,
				Type:  prev.Type,
			}
			vbox.pubsub.Pub(p, pubsub.TopicUpdateSSMValue)

			s := fmt.Sprintf("[green]Value updated: \n%s -> %s", prevValue, newValue)
			vbox.pubsub.Pub(s, pubsub.TopicUpdateInfoBox)
		}

		if prev.Type == ssm.ParameterTypeSecureString {
			vbox.SetText("***************", false)
		}
	}

	return nil
}
