package messages

import (
	"encoding/json"
	"fmt"
	"github.com/notlancer/gpt-cli/internal/builders"
	"github.com/notlancer/gpt-cli/internal/functions"
	"github.com/notlancer/gpt-cli/internal/interfaces"
)

type handlerStruct struct {
	Callback func(MessageProcessor *MessageProcessor, message map[string]interface{}) error
}

var messageProcessorHandlers = map[string]handlerStruct{
	"response.text.delta": {
		Callback: func(_ *MessageProcessor, message map[string]interface{}) error {
			fmt.Print(message["delta"])

			return nil
		},
	},
	"response.content_part.done": {
		Callback: func(messageProcessor *MessageProcessor, message map[string]interface{}) error {
			fmt.Println()

			return messageProcessor.client.StartUserGPTChat()
		},
	},
	"response.function_call_arguments.done": {
		Callback: func(messageProcessor *MessageProcessor, message map[string]interface{}) error {
			return messageProcessor.handleFunctionCall(message)
		},
	},
}

type MessageProcessor struct {
	client interfaces.MessageClient
}

func NewMessageProcessor(client interfaces.MessageClient) *MessageProcessor {
	return &MessageProcessor{client: client}
}

func (p *MessageProcessor) ProcessMessage(message []byte) error {
	var messageData map[string]interface{}
	if err := json.Unmarshal(message, &messageData); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}

	messageType, ok := messageData["type"].(string)
	if !ok {
		return fmt.Errorf("message type is missing or invalid")
	}

	if handler, ok2 := messageProcessorHandlers[messageType]; ok2 {
		return handler.Callback(p, messageData)
	}

	// no message type callback found - ignore
	return nil
}

func (p *MessageProcessor) handleFunctionCall(message map[string]interface{}) error {
	argumentsRaw := message["arguments"].(string)
	var arguments map[string]any

	if err := json.Unmarshal([]byte(argumentsRaw), &arguments); err != nil {
		return fmt.Errorf("failed to unmarshal arguments: %w", err)
	}

	callId := message["call_id"].(string)
	name := message["name"].(string)

	if ok, returnEvent := functions.Handler(name, arguments, callId); ok {
		if err := p.client.SendMessage(returnEvent); err != nil {
			return fmt.Errorf("fail: %w", err)
		}

		responseCreate := builders.BuildResponseCreateMsg()
		if err := p.client.SendMessage(responseCreate); err != nil {
			return fmt.Errorf("fail: %w", err)
		}
	}

	return nil
}
