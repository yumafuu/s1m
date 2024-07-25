package layout

import (
	"github.com/YumaFuu/s1m/tui/cmd"
	"github.com/YumaFuu/s1m/tui/infbox"
	"github.com/YumaFuu/s1m/tui/ptree"
	"github.com/YumaFuu/s1m/tui/vbox"
	"github.com/rivo/tview"
)

type Layout struct {
	*tview.Flex
}

func NewLayout(
	ptree *ptree.ParameterTree,
	infoBox *infbox.InfoBox,
	valueBox *vbox.ValueBox,
	cmdBox *cmd.CmdBox,
) *Layout {
	infvbox := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(infoBox, 0, 1, false).
		AddItem(valueBox, 0, 4, false)

	display := tview.
		NewFlex().
		AddItem(ptree, 0, 1, true).
		AddItem(infvbox, 0, 2, false)

	withCmd := tview.
		NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(display, 0, 1, true).
		AddItem(cmdBox, 2, 1, false)

	return &Layout{withCmd}
}
