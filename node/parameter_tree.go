package component

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func BuildParameterTree(
	parameters map[string]any,
	infoView *tview.TextView,
	valueView *tview.TextView,
) *tview.TreeView {
	root := tview.NewTreeNode(".")

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	addNodes(root, parameters)

	tree.SetBackgroundColor(tcell.ColorDefault)

	return tree
}

func addNodes(parent *tview.TreeNode, m map[string]any) {
	for key, value := range m {
		node := tview.NewTreeNode(key)
		if subMap, ok := value.(map[string]any); ok {
			node.SetColor(tcell.ColorBlue) // Set directories to blue
			node.SetReference(subMap)      // Set reference to the sub-map for later use
			addNodes(node, subMap)
		} else {
			node.SetReference(value) // Set reference to the value for files
		}
		parent.AddChild(node)
	}
}
