package ptree

import (
	"fmt"

	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/atotto/clipboard"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/gdamore/tcell/v2"
)

func (pt *ParameterTree) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	node := pt.GetCurrentNode()
	param, ok := node.GetReference().(types.Parameter)
	if !ok {
		return nil
	}
	clen := len(node.GetChildren())

	switch event.Rune() {
	case 'c':
		if node != nil && clen == 0 {
			var s string
			if err := clipboard.WriteAll(*param.Value); err != nil {
				s = fmt.Sprintf("[red]Error copying to clipboard: %s", err)
			} else {
				s = "[green]Value copied to clipboard"
			}
			pt.pubsub.Pub(s, pubsub.TopicUpdateInfoBox)
		}
	case 'i':
		if node != nil && clen == 0 {
			pt.pubsub.Pub(true, pubsub.TopicSetAppFocusValueBox)
			pt.pubsub.Pub(tcell.ColorBlue, pubsub.TopicUpdateValueBoxBorder)
			pt.pubsub.Pub(param, pubsub.TopicUpdateValueBox)
		}
	case 'q':
		pt.pubsub.Pub(nil, pubsub.TopicStopApp)
	}

	return nil
}
