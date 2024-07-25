package vbox

import (
	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/YumaFuu/ssm-tui/tui/pubsub"
	"github.com/gdamore/tcell/v2"
)

func (vbox *ValueBox) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		vbox.pubsub.Pub(true, pubsub.TopicSetAppFocusTree)

		switch vbox.mode {
		case ModeUpdate:
			prev := vbox.GetPrev()
			prevValue := *prev.Value
			newValue := vbox.GetText()
			if prevValue != newValue {
				p := ssm.Parameter{
					Name:  prev.Name,
					Value: &newValue,
					Type:  prev.Type,
				}
				vbox.pubsub.Pub(p, pubsub.TopicPutSSMValue)
			}

			if prev.Type == ssm.ParameterTypeSecureString {
				vbox.SetText("***************", false)
			}
		case ModeCreate:
			name := vbox.param.Name
			t := vbox.param.Type
			v := vbox.GetText()

			p := ssm.Parameter{
				Name:  name,
				Type:  t,
				Value: &v,
			}
			vbox.pubsub.Pub(p, pubsub.TopicNewParamSubmit)
		}
		vbox.SetBorderColor(tcell.ColorDefault)
	}

	vbox.pubsub.Pub("ESC is pressed", pubsub.TopicAppDraw)
	return event
}
