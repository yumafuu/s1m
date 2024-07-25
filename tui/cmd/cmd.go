package cmd

import (
	"strings"

	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/YumaFuu/ssm-tui/tui/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CmdBox struct {
	*tview.InputField
	pubsub *pubsub.PubSub
}

func NewCmdBox(ps *pubsub.PubSub) *CmdBox {
	style := tcell.StyleDefault.Background(tcell.ColorReset)
	style = style.Attributes(tcell.AttrBold)

	v := tview.NewInputField().
		SetFieldBackgroundColor(tcell.ColorDefault).
		// TODO: ColorDefault is not working
		SetLabelStyle(style).
		SetPlaceholderStyle(tcell.StyleDefault.Foreground(tcell.ColorGray))

	v.
		SetBackgroundColor(tcell.ColorDefault).
		SetBorder(false)

	return &CmdBox{v, ps}
}

func (v *CmdBox) NewParameter(dir string) {
	var param ssm.Parameter

	v.NewParameterValue(dir, param)
}

func (v *CmdBox) NewParameterType(dir string, param ssm.Parameter) {
	v.SetLabel("New Parameter Type: ")
	v.SetPlaceholder("SecureString, String, StringList")
	v.SetAutocompleteFunc(func(currentText string) (entries []string) {
		if len(currentText) == 0 {
			return
		}
		for _, word := range []string{"SecureString", "String", "StringList"} {
			if strings.HasPrefix(strings.ToLower(word), strings.ToLower(currentText)) {
				entries = append(entries, word)
			}
		}
		if len(entries) <= 1 {
			entries = nil
		}
		return
	})
	v.SetAutocompletedFunc(func(text string, index, source int) bool {
		if source != tview.AutocompletedNavigate {
			v.SetText(text)
		}
		return source == tview.AutocompletedEnter || source == tview.AutocompletedClick
	})

	v.SetDoneFunc(func(key tcell.Key) {
		param.Type = ssm.ParameterType(v.GetText())

		if key == tcell.KeyEnter {
			v.NewParameterValue(dir, param)
		}
	})
	return
}

// TODO: Dependency Injection
func (v *CmdBox) NewParameterValue(dir string, param ssm.Parameter) {
	v.SetLabel("New Parameter Name: ")
	v.SetPlaceholder("")
	v.SetText(dir)

	v.SetDoneFunc(func(key tcell.Key) {
		s := v.GetText()
		param.Name = &s

		if key == tcell.KeyEnter {
			v.SetLabel("")
			v.pubsub.Pub(tcell.ColorBlue, pubsub.TopicUpdateValueBoxBorder)
			v.pubsub.Pub(param, pubsub.TopicNewParamCommand)
			v.pubsub.Pub(nil, pubsub.TopicAppDraw)
		}

		v.SetText("")
	})
}

func (v *CmdBox) Confirm(s string, successor func(), finalizer func()) {
	v.SetLabel(s + " (y/n): ")
	v.SetPlaceholder("")
	v.pubsub.Pub(nil, pubsub.TopicAppDraw)

	v.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			if v.GetText() == "y" {
				successor()
			}
		}

		finalizer()
	})
}
