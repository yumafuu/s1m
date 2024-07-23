package ptree

import (
	"strings"

	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/YumaFuu/ssm-tui/aws/ssm"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ParameterTree struct {
	*tview.TreeView
	pubsub pubsub.PubSub
}

const (
	RootNodeName = "."
)

func NewParameterTree(
	pubsub pubsub.PubSub,
	params []ssm.Parameter,
) ParameterTree {
	root := tview.NewTreeNode(RootNodeName)

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)
	tree.SetBackgroundColor(tcell.ColorDefault)

	nodes := buildMapFromPaths(params)
	addNodes(root, nodes)

	pt := ParameterTree{tree, pubsub}
	pt.SetChangedFunc(func(node *tview.TreeNode) {
		pt.displayNodeInfo(node)
	})

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		if len(node.GetChildren()) == 0 {
			pt.displayNodeInfo(node)
		}
		if node.IsExpanded() {
			node.CollapseAll()
		} else {
			node.ExpandAll()
		}
	})

	return pt
}

func addNodes(parent *tview.TreeNode, m map[string]any) {
	for key, value := range m {
		node := tview.NewTreeNode(key)
		if subMap, ok := value.(map[string]any); ok {
			node.SetColor(tcell.ColorBlue)
			node.SetReference(subMap)
			addNodes(node, subMap)
		} else {
			node.SetReference(value)
		}
		parent.AddChild(node)
	}
}

// buildMapFromPaths builds a map from a list of SSM parameters.
// "dir1/dir2/file1", "dir1/dir2/file2" -> { "dir1": { "dir2": { "file1": { ... }, "file2": { ... } } } }
func buildMapFromPaths(params []types.Parameter) map[string]any {
	paramTree := make(map[string]any)
	for _, param := range params {
		key := *param.Name
		parts := strings.Split(strings.TrimPrefix(key, "/"), "/")
		insertPath(paramTree, parts, param)
	}
	return paramTree
}

func insertPath(m map[string]any, parts []string, param types.Parameter) {
	if len(parts) == 0 {
		return
	}
	if len(parts) == 1 {
		m[parts[0]] = param
		return
	}
	if _, ok := m[parts[0]]; !ok {
		m[parts[0]] = make(map[string]any)
	}
	insertPath(m[parts[0]].(map[string]any), parts[1:], param)
}
