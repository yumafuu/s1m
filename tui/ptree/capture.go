package ptree

import (
	"fmt"
	"strings"

	"github.com/YumaFuu/ssm-tui/tui/pubsub"
	"github.com/atotto/clipboard"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (pt *ParameterTree) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	node := pt.GetCurrentNode()
	param, isParam := node.GetReference().(types.Parameter)
	clen := len(node.GetChildren())

	switch event.Rune() {
	case 'c':
		if node != nil && clen == 0 {
			var s string
			if err := clipboard.WriteAll(*param.Value); err != nil {
				s = fmt.Sprintf("[red]Error copying to clipboard: %s", err)
			} else {
				s = "[green]Value is copied to clipboard"
			}
			pt.pubsub.Pub(s, pubsub.TopicUpdateInfoBox)
		}
	case 'y':
		if node != nil && clen == 0 {
			var s string
			if err := clipboard.WriteAll(*param.Name); err != nil {
				s = fmt.Sprintf("[red]Error copying to clipboard: %s", err)
			} else {
				s = fmt.Sprintf("[green]Name `%s` is copied to clipboard", *param.Name)
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
	case 'o':
		var dir string
		if isParam {
			n := *param.Name
			dir = n[:strings.LastIndex(n, "/")] + "/"
		} else {

			var f func(*tview.TreeNode) *types.Parameter
			f = func(n *tview.TreeNode) *types.Parameter {
				for _, p := range n.GetChildren() {
					if len(p.GetChildren()) == 0 {
						if param, ok := p.GetReference().(types.Parameter); ok {
							return &param
						}
					} else {
						f(p)
					}
				}
				return nil
			}

			p := f(node)
			n := *p.Name
			dir = n[:strings.LastIndex(n, "/")] + "/"
		}

		pt.pubsub.Pub(dir, pubsub.TopicNewParam)
	case 'd':
		if node != nil && clen == 0 {
			pt.pubsub.Pub(param, pubsub.TopicDeleteParam)
		}
	case 'r':
		pt.Refresh()
	}

	switch event.Key() {
	case tcell.KeyEnter:
		node.SetExpanded(!node.IsExpanded())
	}

	return event
}
