package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ParameterTree struct {
	Tree *tview.TreeView
}

func (p *ParameterTree) Move(i int) {
	p.Tree.Move(i)
}

func BuildParameterTree(
	parameters map[string]any,
	infoView *tview.TextView,
	valueView *tview.TextArea,
) ParameterTree {
	root := tview.NewTreeNode(".")

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	addNodes(root, parameters)

	tree.SetBackgroundColor(tcell.ColorDefault)

	return ParameterTree{
		Tree: tree,
	}
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
