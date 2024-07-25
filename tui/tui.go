package tui

import (
	"fmt"

	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/YumaFuu/ssm-tui/tui/cmd"
	"github.com/YumaFuu/ssm-tui/tui/infbox"
	"github.com/YumaFuu/ssm-tui/tui/layout"
	"github.com/YumaFuu/ssm-tui/tui/ptree"
	"github.com/YumaFuu/ssm-tui/tui/pubsub"
	"github.com/YumaFuu/ssm-tui/tui/vbox"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type (
	Tui struct {
		app    *tview.Application
		pubsub *pubsub.PubSub
		layout *layout.Layout
		ptree  *ptree.ParameterTree
		infbox *infbox.InfoBox
		vbox   *vbox.ValueBox
		cmdbox *cmd.CmdBox
		ssm    *ssm.Client
	}
)

func NewTui(
	client *ssm.Client,
) (*Tui, error) {
	app := tview.NewApplication()
	app.EnablePaste(true)

	ps := pubsub.NewPubSub()

	pst, err := ptree.NewParameterTree(
		ps,
		client,
	)
	if err != nil {
		return nil, err
	}

	infbox := infbox.NewInfoBox(ps)
	vbox := vbox.NewValueBox(ps)
	cmdbox := cmd.NewCmdBox(ps)

	layout := layout.NewLayout(pst, infbox, vbox, cmdbox)
	app.SetRoot(layout, true)

	a := &Tui{
		app:    app,
		pubsub: ps,
		layout: layout,
		ptree:  pst,
		infbox: infbox,
		vbox:   vbox,
		cmdbox: cmdbox,
		ssm:    client,
	}

	a.SetInputCapture()
	return a, nil
}

func (a *Tui) Run() error {
	go a.infbox.WaitTopic()
	go a.vbox.WaitTopic()
	go a.WaitTopic()

	if err := a.app.Run(); err != nil {
		panic(err)
	}
	return nil
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
						a.cmdbox.SetLabel("")
						a.cmdbox.SetText("")
						a.ptree.Refresh()
					}

					a.app.SetFocus(a.ptree)
				})
			// if err := a.ssm.Put(
			// 	param.Name,
			// 	param.Type,
			// 	param.Value,
			// ); err != nil {
			// 	a.infbox.SetText(err.Error())
			// }

			// a.ptree.Refresh()
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
		case inp := <-chNewParamSubmit:
			param, ok := inp.(ssm.Parameter)
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
						a.cmdbox.SetLabel("")
						a.cmdbox.SetText("")
						a.ptree.Refresh()
					}

					a.app.SetFocus(a.ptree)
				})
		}
	}
}
