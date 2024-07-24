package tui

import (
	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/YumaFuu/ssm-tui/tui/infbox"
	"github.com/YumaFuu/ssm-tui/tui/layout"
	"github.com/YumaFuu/ssm-tui/tui/ptree"
	"github.com/YumaFuu/ssm-tui/tui/pubsub"
	"github.com/YumaFuu/ssm-tui/tui/vbox"
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

	layout := layout.NewLayout(pst, infbox, vbox)
	app.SetRoot(layout, true)

	a := &Tui{
		app:    app,
		pubsub: ps,
		layout: layout,
		ptree:  pst,
		infbox: infbox,
		vbox:   vbox,
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
	chDraw := a.pubsub.Sub(pubsub.TopicAppDraw)
	chUpdateSSMValue := a.pubsub.Sub(pubsub.TopicUpdateSSMValue)

	for {
		select {
		case <-chStop:
			a.app.Stop()
		case <-chFocusTree:
			a.app.SetFocus(a.ptree)
		case <-chFocusVBox:
			a.app.SetFocus(a.vbox)
		case <-chDraw:
			a.app.Draw()
		case msg := <-chUpdateSSMValue:
			param, ok := msg.(ssm.Parameter)
			if !ok {
				continue
			}

			if err := a.ssm.Update(
				param.Name,
				param.Type,
				param.Value,
			); err != nil {
				a.infbox.SetText(err.Error())
			}

			a.ptree.Refresh()
		}
	}
}
