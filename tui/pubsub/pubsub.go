package pubsub

import "github.com/cskr/pubsub/v2"

type (
	PubSub = pubsub.PubSub[string, any]
)

const (
	TopicUpdateInfoBox        = "UpdateInfoBox"
	TopicUpdateValueBox       = "UpdateValueBox"
	TopicUpdateValueBoxBorder = "UpdateValueBoxBorder"
	TopicSetAppFocusValueBox  = "SetFocusValueBox"
	TopicStopApp              = "Stop"
	TopicSetAppFocusTree      = "SetFocusTree"
	TopicAppDraw              = "AppDraw"
	TopicPutSSMValue          = "PutSSMValue"
	TopicCreateSSMValue       = "CreateSSMValue"
	TopicDeleteParam          = "DeleteParam"
	TopicNewParam             = "NewParam"
	TopicNewParamCommand      = "NewParamCommand"
	TopicNewParamSubmit       = "NewParamSubmit"
	TopicPtreeRefresh         = "PtreeRefresh"

	CAPACITY = 30
)

func NewPubSub() *PubSub {
	return pubsub.New[string, any](CAPACITY)
}
