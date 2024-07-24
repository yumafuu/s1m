package pubsub

import "github.com/cskr/pubsub/v2"

type (
	PubSub = *pubsub.PubSub[string, any]
)

const (
	TopicUpdateInfoBox        = "UpdateInfoBox"
	TopicUpdateValueBox       = "UpdateValueBox"
	TopicUpdateValueBoxBorder = "UpdateValueBoxBorder"
	TopicSetAppFocusValueBox  = "SetFocusValueBox"
	TopicStopApp              = "Stop"
	TopicSetAppFocusTree      = "SetFocusTree"
	TopicAppDraw              = "AppDraw"
)

func NewPubSub() PubSub {
	return pubsub.New[string, any](0)
}
