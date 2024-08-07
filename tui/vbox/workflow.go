package vbox

import (
	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui/pubsub"
	"github.com/gdamore/tcell/v2"
)

func (v *ValueBox) WorkflowUpdateParam(param ssm.Parameter) {
	v.mode = ModeUpdate
	v.SetPrev(param)

	v.Input.SetBorderColor(tcell.ColorBlue)
	v.TextArea.SetBorderColor(tcell.ColorBlue)
	v.ListView.SetBorderColor(tcell.ColorBlue)

	v.ShowPage(KeyTextArea)

	v.TextArea.SetTitle("Value")
	v.TextArea.SetText(*param.Value, true)
	v.TextArea.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		s := v.GetText()
		prev := v.GetPrev().Value

		v.param.Value = &s

		if event.Key() == tcell.KeyESC {
			v.TextArea.SetBorderColor(tcell.ColorDefault)
			v.TextArea.SetTitle("")
			v.pubsub.Pub(nil, pubsub.TopicAppDraw)
			if *prev != s {
				v.pubsub.Pub(v.param, pubsub.TopicUpdateParamSubmit)
			} else {
				v.pubsub.Pub(nil, pubsub.TopicFocusTree)
			}
		}
		return event
	})
}

func (v *ValueBox) WorkflowCreateParam(dir string) {
	v.mode = ModeCreate

	v.param = ssm.Parameter{}
	v.pubsub.Pub(v.param, pubsub.TopicWriteInfoBox)

	v.Input.SetBorderColor(tcell.ColorGreen)
	v.TextArea.SetBorderColor(tcell.ColorGreen)
	v.ListView.SetBorderColor(tcell.ColorGreen)

	v.ShowPage(KeyInput)
	v.Input.SetText(dir)
	v.Input.SetTitle("Name")
	v.pubsub.Pub(dir, pubsub.TopicAppDraw)

	v.Input.SetDoneFunc(func(key tcell.Key) {
		s := v.Input.GetText()
		v.param.Name = &s
		v.pubsub.Pub(v.param, pubsub.TopicWriteInfoBox)

		v.HidePage(KeyInput)
		v.ShowPage(KeyList)
	})

	typeList := []ssm.ParameterType{
		ssm.ParameterTypeString,
		ssm.ParameterTypeSecureString,
		ssm.ParameterTypeStringList,
	}

	v.ListView.Clear()
	v.ListView.SetTitle("Select Type")
	v.ListView.
		AddItem(string(ssm.ParameterTypeString), "", '1', nil).
		AddItem(string(ssm.ParameterTypeSecureString), "", '2', nil).
		AddItem(string(ssm.ParameterTypeStringList), "", '3', nil).
		SetSelectedFunc(func(i int, _ string, _ string, _ rune) {
			v.param.Type = typeList[i]
			v.pubsub.Pub(v.param, pubsub.TopicWriteInfoBox)

			v.HidePage(KeyList)
			v.ShowPage(KeyTextArea)
		})

	v.TextArea.SetTitle("Value")
	v.TextArea.SetText("", true)
	v.TextArea.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		s := v.GetText()
		v.param.Value = &s

		if event.Key() == tcell.KeyESC {
			v.TextArea.SetBorderColor(tcell.ColorDefault)

			v.TextArea.SetTitle("")
			v.pubsub.Pub(nil, pubsub.TopicAppDraw)
			v.pubsub.Pub(v.param, pubsub.TopicCreateParamSubmit)
		}
		return event
	})

}
