package app

import (
	"github.com/YumaFuu/ssm-tui/app/infbox"
	"github.com/YumaFuu/ssm-tui/app/ptree"
	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/YumaFuu/ssm-tui/app/vbox"
	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/rivo/tview"
)

type (
	App struct {
		*tview.Application
		pubsub pubsub.PubSub
		layout Layout
		ptree  ptree.ParameterTree
		infbox infbox.InfoBox
		vbox   vbox.ValueBox
	}
)

func NewApp(params []ssm.Parameter) *App {
	app := tview.NewApplication()

	ps := pubsub.NewPubSub()

	pst := ptree.NewParameterTree(ps, params)
	infbox := infbox.NewInfoBox(ps)
	vbox := vbox.NewValueBox(ps)

	layout := NewLayout(pst, infbox, vbox)

	a := &App{
		Application: app,
		pubsub:      ps,
		layout:      layout,
		ptree:       pst,
		infbox:      infbox,
		vbox:        vbox,
	}

	a.SetInputCapture(a.InputCapture)
	return a
}

func (a *App) Run() error {
	go a.infbox.WaitTopic()
	go a.vbox.WaitTopic()
	go a.WaitTopic()

	if err := a.SetRoot(a.layout, true).Run(); err != nil {
		panic(err)
	}
	return nil
}

func (a *App) WaitTopic() {
	chStop := a.pubsub.Sub(pubsub.TopicStopApp)
	chFocusTree := a.pubsub.Sub(pubsub.TopicSetAppFocusTree)
	chFocusVBox := a.pubsub.Sub(pubsub.TopicSetAppFocusValueBox)

	for {
		select {
		case <-chStop:
			a.Stop()
		case <-chFocusTree:
			a.SetFocus(a.ptree)
		case <-chFocusVBox:
			a.SetFocus(a.vbox)
		}
	}
}
