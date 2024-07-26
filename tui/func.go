package tui

import (
	"fmt"

	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui/cmd"
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

func (a *Tui) Reload() error {
	a.app.Draw()
	if err := a.ptree.Refresh(); err != nil {
		return err
	}
	a.app.SetFocus(a.ptree)
	a.app.Draw()

	return nil
}

func (a *Tui) WaitTopic() {
	chReload := a.pubsub.Sub(pubsub.TopicAppReload)
	chFocusTree := a.pubsub.Sub(pubsub.TopicFocusTree)
	chFocusVBox := a.pubsub.Sub(pubsub.TopicFocusValueBox)
	chUpdateParamStart := a.pubsub.Sub(pubsub.TopicUpdateParamStart)
	chUpdateParamSubmit := a.pubsub.Sub(pubsub.TopicUpdateParamSubmit)
	chCreateParamStart := a.pubsub.Sub(pubsub.TopicCreateParamStart)
	chCreateParamSubmit := a.pubsub.Sub(pubsub.TopicCreateParamSubmit)
	chDeleteParam := a.pubsub.Sub(pubsub.TopicDeleteParam)

	for {
		select {
		case <-chReload:
			if err := a.Reload(); err != nil {
				a.infbox.SetText(err.Error())
			}
		case <-chFocusTree:
			a.app.SetFocus(a.ptree)
		case <-chFocusVBox:
			a.vbox.SetMode(vbox.ModeUpdate)
			a.app.SetFocus(a.vbox)
		case msg := <-chUpdateParamStart:
			param, ok := msg.(ssm.Parameter)
			if !ok {
				continue
			}

			a.vbox.WorkflowUpdateParam(param)
			a.app.SetFocus(a.vbox)
			a.app.Draw()
		case msg := <-chUpdateParamSubmit:
			param, ok := msg.(ssm.Parameter)
			if !ok {
				continue
			}

			a.app.SetFocus(a.cmdbox)
			a.cmdbox.Confirm(cmd.ConfirmInput{
				Label: "[blue]Are you sure to Update?",
				Successor: func() {
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
					a.vbox.SetBorderColor(tcell.ColorDefault)
				},
			})

		case dir := <-chCreateParamStart:
			s, ok := dir.(string)
			if !ok {
				continue
			}

			a.vbox.WorkflowCreateParam(s)
			a.app.SetFocus(a.vbox)

		case msg := <-chCreateParamSubmit:
			param, ok := msg.(ssm.Parameter)
			if !ok {
				continue
			}

			a.app.SetFocus(a.cmdbox)
			a.cmdbox.Confirm(cmd.ConfirmInput{
				Label: "[green]Are you sure to Create?",
				Successor: func() {
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
					a.vbox.SetBorderColor(tcell.ColorDefault)
				},
			})
		case msg := <-chDeleteParam:
			param, ok := msg.(ssm.Parameter)
			if !ok {
				continue
			}

			a.app.SetFocus(a.cmdbox)
			a.cmdbox.Confirm(
				cmd.ConfirmInput{
					Label: "[red]Are you sure to Delete?",
					Successor: func() {
						if err := a.ssm.Delete(param.Name); err != nil {
							a.infbox.SetText(err.Error())
						} else {
							a.infbox.SetText(fmt.Sprintf(
								infbox.DeleteMessageFormat,
								*param.Name,
							))
						}
						a.vbox.SetBorderColor(tcell.ColorDefault)
					},
				},
			)
			a.app.Draw()
		}
	}
}
