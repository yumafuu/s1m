package app

import (
	"github.com/YumaFuu/ssm-tui/app/infbox"
	"github.com/YumaFuu/ssm-tui/app/ptree"
	"github.com/YumaFuu/ssm-tui/app/vbox"
	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/cskr/pubsub/v2"
	"github.com/rivo/tview"
)

type (
	App struct {
		app    *tview.Application
		pubsub *pubsub.PubSub[string, string]
		layout Layout
	}
)

func NewApp(params []ssm.Parameter) *App {
	app := tview.NewApplication()
	ps := pubsub.New[string, string](0)

	pst := ptree.NewParameterTree(ps, params)
	info := infbox.NewInfoBox(ps)
	value := vbox.NewValueBox(ps)

	go info.WaitTopic()
	go value.WaitTopic()

	layout := NewLayout(pst, info, value)

	return &App{
		app:    app,
		pubsub: ps,
		layout: layout,
	}
}

func (a *App) Run() error {
	if err := a.app.SetRoot(a.layout, true).Run(); err != nil {
		panic(err)
	}
	return nil
}
