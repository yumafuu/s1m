package component

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/rivo/tview"
)

func DisplayNodeInfo(
	node *tview.TreeNode,
	infoView *tview.TextView,
) {
	if node == nil || len(node.GetChildren()) != 0 {
		infoView.SetText("")
		return
	}
	param := node.GetReference().(types.Parameter)
	info := fmt.Sprintf("Name: %s\nType: %s\nLastModifiedDate: %s\nVersion: %d\n\n%s\n\n",
		*param.Name, param.Type, param.LastModifiedDate, param.Version, *param.Value)
	infoView.SetText(info)
}
