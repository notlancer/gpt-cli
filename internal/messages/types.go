package messages

type MessageProcessorHandlerCallback func(MessageProcessor *MessageProcessor, message map[string]interface{}) error

type MessageProcessorHandler struct {
	Callback MessageProcessorHandlerCallback
}
