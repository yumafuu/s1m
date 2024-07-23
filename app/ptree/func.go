package ptree

import (
	"fmt"

	"github.com/YumaFuu/ssm-tui/app/pubsub"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/rivo/tview"
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
	param := node.GetReference().(types.Parameter)
	info := fmt.Sprintf(
		`Version:          %d
Name:             %s
Type:             %s
LastModifiedDate: %s`,
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

	go pt.pubsub.Pub(s, pubsub.TopicUpdateValueBox)
	go pt.pubsub.Pub(info, pubsub.TopicUpdateInfoBox)
}
