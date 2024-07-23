package pubsub

import "github.com/cskr/pubsub/v2"

type (
	PubSub = *pubsub.PubSub[string, string]
)

const (
	TopicUpdateInfoBox  = "UpdateInfoBox"
	TopicUpdateValueBox = "UpdateValueBox"
)
