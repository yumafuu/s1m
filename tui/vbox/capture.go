package vbox

import (
	"fmt"

	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui/pubsub"
	"github.com/gdamore/tcell/v2"
)

func (v *ValueBox) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		switch v.mode {
		case ModeUpdate:
			prev := v.GetPrev()
			prevValue := *prev.Value
			newValue := v.GetText()
			v.pubsub.Pub(fmt.Sprintf("%s -> %s", prevValue, newValue), pubsub.TopicWriteInfoBox)
			fmt.Println("ESC is pressed")

			if prevValue != newValue {
				p := ssm.Parameter{
					Name:  prev.Name,
					Value: &newValue,
					Type:  prev.Type,
				}
				v.pubsub.Pub(p, pubsub.TopicUpdateParamSubmit)
			}

			if prev.Type == ssm.ParameterTypeSecureString {
				v.SetText("***************", false)
			}
		case ModeCreate:
			name := v.param.Name
			t := v.param.Type
			vl := v.GetText()

			p := ssm.Parameter{
				Name:  name,
				Type:  t,
				Value: &vl,
			}
			v.pubsub.Pub(p, pubsub.TopicCreateParamSubmit)
		}
	}

	v.pubsub.Pub("ESC is pressed", pubsub.TopicAppDraw)
	return event
}
