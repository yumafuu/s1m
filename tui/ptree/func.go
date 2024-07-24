package ptree

import (
	"fmt"
	"sort"
	"strings"

	"github.com/YumaFuu/ssm-tui/tui/pubsub"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	infoFormat = `Version:          %d
Name:             %s
Type:             %s
LastModifiedDate: %s`
)

// // Function to display information of the selected node
func (pt *ParameterTree) displayNodeInfo(node *tview.TreeNode) {
	if node == nil {
		pt.pubsub.Pub("", pubsub.TopicUpdateInfoBox)
		pt.pubsub.Pub("", pubsub.TopicUpdateValueBox)
		return
	}
	if len(node.GetChildren()) != 0 {
		pt.pubsub.Pub("", pubsub.TopicUpdateInfoBox)
		pt.pubsub.Pub("", pubsub.TopicUpdateValueBox)
		return
	}
	param, ok := node.GetReference().(types.Parameter)
	if !ok {
		pt.pubsub.Pub("", pubsub.TopicUpdateInfoBox)
		pt.pubsub.Pub("", pubsub.TopicUpdateValueBox)
		return
	}

	info := fmt.Sprintf(
		infoFormat,
		param.Version,
		*param.Name,
		param.Type,
		param.LastModifiedDate,
	)
	var s string
	if param.Type == types.ParameterTypeSecureString {
		s = "***************"
	} else {
		s = *param.Value
	}

	param.Value = &s

	go pt.pubsub.Pub(param, pubsub.TopicUpdateValueBox)
	go pt.pubsub.Pub(info, pubsub.TopicUpdateInfoBox)
}

func (pt *ParameterTree) Refresh() error {
	var currentName string

	p, ok := pt.GetCurrentNode().GetReference().(types.Parameter)
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
			if param, ok := node.GetReference().(types.Parameter); ok {
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

func addNodes(parent *tview.TreeNode, m map[string]any) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// for key, value := range m {
	for _, key := range keys {
		value := m[key]
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
