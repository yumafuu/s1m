package app

import (
	"github.com/YumaFuu/ssm-tui/app/infbox"
	"github.com/YumaFuu/ssm-tui/app/layout"
	"github.com/YumaFuu/ssm-tui/app/ptree"
	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/YumaFuu/ssm-tui/app/vbox"
	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/rivo/tview"
)

type (
	App struct {
		tapp   *tview.Application
		pubsub *pubsub.PubSub
		layout *layout.Layout
		ptree  *ptree.ParameterTree
		infbox *infbox.InfoBox
		vbox   *vbox.ValueBox
		ssm    *ssm.Client
	}
)

func NewApp(
	client *ssm.Client,
) (*App, error) {
	tapp := tview.NewApplication()
	tapp.EnablePaste(true)

	ps := pubsub.NewPubSub()

	pst, err := ptree.NewParameterTree(ps, client)
	if err != nil {
		return nil, err
	}

	infbox := infbox.NewInfoBox(ps)
	vbox := vbox.NewValueBox(ps)

	layout := layout.NewLayout(pst, infbox, vbox)
	tapp.SetRoot(layout, true)

	a := &App{
		tapp:   tapp,
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

func (a *App) Run() error {
	go a.infbox.WaitTopic()
	go a.vbox.WaitTopic()
	go a.WaitTopic()

	if err := a.tapp.Run(); err != nil {
		panic(err)
	}
	return nil
}

func (a *App) WaitTopic() {
	chStop := a.pubsub.Sub(pubsub.TopicStopApp)
	chFocusTree := a.pubsub.Sub(pubsub.TopicSetAppFocusTree)
	chFocusVBox := a.pubsub.Sub(pubsub.TopicSetAppFocusValueBox)
	chDraw := a.pubsub.Sub(pubsub.TopicAppDraw)
	chUpdateSSMValue := a.pubsub.Sub(pubsub.TopicUpdateSSMValue)

	for {
		select {
		case <-chStop:
			a.tapp.Stop()
		case <-chFocusTree:
			a.tapp.SetFocus(a.ptree)
		case <-chFocusVBox:
			a.tapp.SetFocus(a.vbox)
		case <-chDraw:
			a.tapp.Draw()
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
