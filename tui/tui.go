package tui

import (
	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/YumaFuu/ssm-tui/tui/cmd"
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
