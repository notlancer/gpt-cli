package messages

import "github.com/notlancer/gpt-cli/internal/interfaces"

type MessageProcessorHandlerCallback func(MessageProcessor *MessageProcessor, message map[string]interface{}) error

type MessageProcessorHandler struct {
	Callback MessageProcessorHandlerCallback
}

type MessageProcessor struct {
	client interfaces.MessageClient
}
