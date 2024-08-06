package ptree

import (
	"fmt"
	"sort"
	"strings"

	"github.com/YumaFuu/s1m/aws/ssm"
	"github.com/YumaFuu/s1m/tui/infbox"
	"github.com/YumaFuu/s1m/tui/pubsub"
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

// // Function to display information of the selected node
func (pt *ParameterTree) displayNodeInfo(node *tview.TreeNode) {
	if node == nil || len(node.GetChildren()) != 0 {
		pt.pubsub.Pub("", pubsub.TopicWriteInfoBox)
		pt.pubsub.Pub("", pubsub.TopicWriteValueBox)
		return
	}

	param, ok := node.GetReference().(ssm.Parameter)
	if !ok {
		pt.pubsub.Pub("", pubsub.TopicWriteInfoBox)
		pt.pubsub.Pub("", pubsub.TopicWriteValueBox)
		return
	}

	info := fmt.Sprintf(
		infbox.ValueFormat,
		param.Version,
		*param.Name,
		param.Type,
		param.LastModifiedDate,
	)
	var s string
	if param.Type == ssm.ParameterTypeSecureString {
		s = "***************"
	} else {
		s = *param.Value
	}

	param.Value = &s

	go pt.pubsub.Pub(param, pubsub.TopicWriteValueBox)
	go pt.pubsub.Pub(info, pubsub.TopicWriteInfoBox)
}

func (pt *ParameterTree) Refresh() error {
	var currentName string

	p, ok := pt.GetCurrentNode().GetReference().(ssm.Parameter)
	if ok {
		currentName = *p.Name
	}

	pt.root.ClearChildren()

	params, err := pt.client.List("/")
	if err != nil {
		return err
	}

	nodes := buildMapFromPaths(params)
	addNodes(pt.root, nodes)

	pt.SetChangedFunc(func(node *tview.TreeNode) {
		pt.displayNodeInfo(node)
	})

	pt.SetSelectedFunc(func(node *tview.TreeNode) {
		if len(node.GetChildren()) == 0 {
			pt.displayNodeInfo(node)
		}
		if node.IsExpanded() {
			node.CollapseAll()
		} else {
			node.ExpandAll()
		}
	})

	var setCurrentNode func(node *tview.TreeNode)
	setCurrentNode = func(node *tview.TreeNode) {
		if node == nil {
			return
		}
		if len(node.GetChildren()) == 0 {
			if param, ok := node.GetReference().(ssm.Parameter); ok {
				if *param.Name == currentName {
					pt.SetCurrentNode(node)
				}
			}
		} else {
			for _, child := range node.GetChildren() {
				setCurrentNode(child)
			}
		}
	}

	for _, node := range pt.root.GetChildren() {
		setCurrentNode(node)
	}

	return nil
}

func addNodes(parent *tview.TreeNode, m Node) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := m[key]
		node := tview.NewTreeNode(key)
		if subMap, ok := value.(Node); ok {
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
func buildMapFromPaths(params []ssm.Parameter) Node {
	paramTree := make(Node)
	var keys []int
	for k := range params {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		param := params[k]

		key := *param.Name
		parts := strings.Split(strings.TrimPrefix(key, "/"), "/")
		insertPath(paramTree, parts, param)
	}

	return paramTree
}

func insertPath(m Node, parts []string, param ssm.Parameter) {
	if len(parts) == 0 {
		return
	}
	if len(parts) == 1 {
		m[parts[0]] = param
		return
	}

	if _, ok := m[parts[0]]; !ok {
		m[parts[0]] = make(Node)
	}

	if mm, ok := m[parts[0]].(Node); ok {
		insertPath(mm, parts[1:], param)
	}
}
