package vbox

import (
	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/gdamore/tcell/v2"
)

func (vbox *ValueBox) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	switch event.Key() {
	case tcell.KeyEsc:
		vbox.pubsub.Pub(true, pubsub.TopicSetAppFocusTree)
		vbox.SetBorderColor(tcell.ColorDefault)
		vbox.SetBorderStyle(tcell.StyleDefault.Blink(false))

		// node := tree.GetCurrentNode()
		// prev := *node.GetReference().(types.Parameter).Value
		// new := valueView.GetText()
		// if prev != new {
		// 	layout.ShowPage("confirm")
		// 	layout.HidePage("confirm")
		// 	infoView.SetText(
		// 		fmt.Sprintf("[green]Value updated: \n%s -> %s", prev, new),
		// 	)
		// }
	}

	return nil
}
