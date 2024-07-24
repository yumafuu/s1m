package vbox

import (
	"fmt"

	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/gdamore/tcell/v2"
)

func (vbox *ValueBox) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		vbox.pubsub.Pub(true, pubsub.TopicSetAppFocusTree)
		vbox.pubsub.Pub("ESC is pressed", pubsub.TopicAppDraw)
		vbox.SetBorderColor(tcell.ColorDefault)

		prev := vbox.GetPrev()
		new := vbox.GetText()
		if prev != new {
			s := fmt.Sprintf("[green]Value updated: \n%s -> %s", prev, new)
			vbox.pubsub.Pub(s, pubsub.TopicUpdateInfoBox)
		}
	}

	return nil
}
