package ptree

import (
	"fmt"
	"strings"

	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui/pubsub"
	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (pt *ParameterTree) InputCapture(event *tcell.EventKey) *tcell.EventKey {
	currentNode := pt.GetCurrentNode()
	param, isParam := currentNode.GetReference().(ssm.Parameter)
	clen := len(currentNode.GetChildren())

	switch event.Rune() {
	case 'c':
		if currentNode != nil && clen == 0 {
			var s string
			if err := clipboard.WriteAll(*param.Value); err != nil {
				s = fmt.Sprintf("[red]Error copying to clipboard: %s", err)
			} else {
				s = "[green]Value is copied to clipboard"
			}
			pt.pubsub.Pub(s, pubsub.TopicWriteInfoBox)
		}
	case 'y':
		if currentNode != nil && clen == 0 {
			var s string
			if err := clipboard.WriteAll(*param.Name); err != nil {
				s = fmt.Sprintf("[red]Error copying to clipboard: %s", err)
			} else {
				s = fmt.Sprintf("[green]Name `%s` is copied to clipboard", *param.Name)
			}
			pt.pubsub.Pub(s, pubsub.TopicWriteInfoBox)
		}

	case 'i':
		if currentNode != nil && clen == 0 {
			param, ok := currentNode.GetReference().(ssm.Parameter)
			if !ok {
				break
			}

			pt.pubsub.Pub(param, pubsub.TopicUpdateParamStart)
			pt.pubsub.Pub(nil, pubsub.TopicAppDraw)
		}
	case 'o':
		var dir string
		if isParam {
			n := *param.Name
			dir = n[:strings.LastIndex(n, "/")] + "/"
		} else {

			var f func(*tview.TreeNode) *ssm.Parameter
			f = func(n *tview.TreeNode) *ssm.Parameter {
				for _, p := range n.GetChildren() {
					if len(p.GetChildren()) == 0 {
						if param, ok := p.GetReference().(ssm.Parameter); ok {
							return &param
						}
					} else {
						f(p)
					}
				}
				return nil
			}

			p := f(currentNode)
			n := *p.Name
			dir = n[:strings.LastIndex(n, "/")] + "/"
		}

		pt.pubsub.Pub(dir, pubsub.TopicCreateParamStart)
	case 'd':
		if currentNode != nil && clen == 0 {
			pt.pubsub.Pub(param, pubsub.TopicDeleteParam)
			pt.pubsub.Pub(nil, pubsub.TopicAppDraw)
		}
	case 'r':
		if err := pt.Refresh(); err != nil {
			pt.pubsub.Pub(err.Error(), pubsub.TopicWriteInfoBox)
		}
	}

	return event
}
