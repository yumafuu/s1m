package ptree

import (
	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/YumaFuu/ssm-tui/tui/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type (
	ParameterTree struct {
		*tview.TreeView
		pubsub   *pubsub.PubSub
		client   *ssm.Client
		root     *tview.TreeNode
		position int
	}
	Node = map[string]any
)

const (
	ROOT_NODENAME = "."
	INIT_POSITION = 0
)

func NewParameterTree(
	pubsub *pubsub.PubSub,
	client *ssm.Client,
) (*ParameterTree, error) {
	root := tview.NewTreeNode(ROOT_NODENAME)

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)
	tree.SetBackgroundColor(tcell.ColorDefault)

	pt := &ParameterTree{
		tree,
		pubsub,
		client,
		root,
		INIT_POSITION,
	}
	if err := pt.Refresh(); err != nil {
		return nil, err
	}

	return pt, nil
}
