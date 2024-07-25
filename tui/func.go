package tui

import (
	"fmt"

	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui/infbox"
	"github.com/YumaFuu/s1m/tui/pubsub"
	"github.com/YumaFuu/s1m/tui/vbox"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (a *Tui) SetInputCapture() *tview.Application {
	inputCapture := func(event *tcell.EventKey) *tcell.EventKey {
		switch a.app.GetFocus() {
		case a.ptree:
			a.ptree.InputCapture(event)
		case a.vbox:
			a.vbox.InputCapture(event)
		}

		if event.Key() == tcell.KeyCtrlC {
			a.app.Stop()
		}

		return event
	}
	a.app.SetInputCapture(inputCapture)

	return a.app
}

func (a *Tui) WaitTopic() {
	chStop := a.pubsub.Sub(pubsub.TopicStopApp)
	chFocusTree := a.pubsub.Sub(pubsub.TopicSetAppFocusTree)
	chFocusVBox := a.pubsub.Sub(pubsub.TopicSetAppFocusValueBox)
	chUpdateVBoxBorder := a.pubsub.Sub(pubsub.TopicUpdateValueBoxBorder)
	chDraw := a.pubsub.Sub(pubsub.TopicAppDraw)
	chUpdateSSMValue := a.pubsub.Sub(pubsub.TopicPutSSMValue)
	chNewParam := a.pubsub.Sub(pubsub.TopicNewParam)
	chNewParamCommand := a.pubsub.Sub(pubsub.TopicNewParamCommand)
	chNewParamSubmit := a.pubsub.Sub(pubsub.TopicNewParamSubmit)
	chDeleteParam := a.pubsub.Sub(pubsub.TopicDeleteParam)

	for {
		select {
		case <-chStop:
			a.app.Stop()
		case <-chFocusTree:
			a.app.SetFocus(a.ptree)
		case <-chFocusVBox:
			a.vbox.SetMode(vbox.ModeUpdate)
			a.app.SetFocus(a.vbox)
		case <-chDraw:
			a.app.Draw()
		case msg := <-chUpdateSSMValue:
			param, ok := msg.(ssm.Parameter)
			if !ok {
				continue
			}

			a.app.SetFocus(a.cmdbox)
			a.app.Draw()
			a.cmdbox.Confirm(
				"Are you sure to Update?",
				func() {
					if err := a.ssm.Put(
						param.Name,
						param.Type,
						param.Value,
					); err != nil {
						a.infbox.SetText(err.Error())
					} else {
						a.infbox.SetText(fmt.Sprintf(
							infbox.UpdateMessageFormat,
							*param.Name,
							param.Type,
							*param.Value,
						))
					}
				},
				func() {
					a.cmdbox.SetLabel("")
					a.cmdbox.SetText("")
					if err := a.ptree.Refresh(); err != nil {
						a.infbox.SetText(err.Error())
					}
					a.app.SetFocus(a.ptree)
				},
			)
		case color := <-chUpdateVBoxBorder:
			c, ok := color.(tcell.Color)
			if !ok {
				continue
			}
			a.vbox.SetBorderColor(c)
		case dir := <-chNewParam:
			s, ok := dir.(string)
			if !ok {
				continue
			}

			a.cmdbox.NewParameter(s)
			a.app.SetFocus(a.cmdbox)
		case inp := <-chNewParamCommand:
			param, ok := inp.(ssm.Parameter)
			if !ok {
				continue
			}

			a.infbox.SetText(fmt.Sprintf(
				infbox.ValueFormat,
				0,
				*param.Name,
				param.Type,
				"",
			))
			a.vbox.SetText("", true)
			a.vbox.SetMode(vbox.ModeCreate)
			a.vbox.SetParam(param)

			a.app.SetFocus(a.vbox)
		case msg := <-chNewParamSubmit:
			param, ok := msg.(ssm.Parameter)
			if !ok {
				continue
			}

			a.app.SetFocus(a.cmdbox)
			a.cmdbox.Confirm(
				"Are you sure to Create?",
				func() {
					if err := a.ssm.Put(
						param.Name,
						param.Type,
						param.Value,
					); err != nil {
						a.infbox.SetText(err.Error())
					} else {
						a.infbox.SetText(fmt.Sprintf(
							infbox.CreateMessageFormat,
							*param.Name,
							param.Type,
							*param.Value,
						))
					}
				},
				func() {
					a.cmdbox.SetLabel("")
					a.cmdbox.SetText("")
					if err := a.ptree.Refresh(); err != nil {
						a.infbox.SetText(err.Error())
					}
					a.app.SetFocus(a.ptree)
				},
			)
		case msg := <-chDeleteParam:
			param, ok := msg.(ssm.Parameter)
			if !ok {
				continue
			}
			a.app.SetFocus(a.cmdbox)
			a.cmdbox.Confirm(
				"Are you sure to Delete?",
				func() {
					if err := a.ssm.Delete(
						param.Name,
					); err != nil {
						a.infbox.SetText(err.Error())
					} else {
						a.infbox.SetText(fmt.Sprintf(
							infbox.DeleteMessageFormat,
							*param.Name,
						))
					}
				},
				func() {
					a.cmdbox.SetLabel("")
					a.cmdbox.SetText("")
					if err := a.ptree.Refresh(); err != nil {
						a.infbox.SetText(err.Error())
					}
					a.app.SetFocus(a.ptree)
				},
			)
		}
	}
}
