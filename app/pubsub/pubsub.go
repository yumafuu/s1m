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
	TopicUpdateSSMValue       = "UpdateSSMValue"

	CAPACITY = 10
)

func NewPubSub() *PubSub {
	return pubsub.New[string, any](CAPACITY)
}
