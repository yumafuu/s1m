package pubsub

import "github.com/cskr/pubsub/v2"

type (
	PubSub = pubsub.PubSub[string, any]
)

const (
	TopicWriteInfoBox      = "WriteInfoBox"
	TopicWriteValueBox     = "WriteValueBox"
	TopicFocusValueBox     = "SetFocusValueBox"
	TopicFocusTree         = "SetFocusTree"
	TopicAppReload         = "AppReload"
	TopicAppDraw           = "AppDraw"
	TopicCreateSSMValue    = "CreateSSMValue"
	TopicUpdateParamStart  = "UpdateParamStart"
	TopicUpdateParamSubmit = "UpdateParamSubmit"
	TopicCreateParamStart  = "CreateParamStart"
	TopicCreateParamSubmit = "CreateParamSubmit"
	TopicDeleteParam       = "DeleteParam"

	CAPACITY = 30
)

func NewPubSub() *PubSub {
	return pubsub.New[string, any](CAPACITY)
}
