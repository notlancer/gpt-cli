package messages

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/notlancer/gpt-cli/internal/builders"
	"github.com/notlancer/gpt-cli/internal/functions"
	"github.com/notlancer/gpt-cli/internal/interfaces"
	"github.com/notlancer/gpt-cli/internal/validation"
)

var messageProcessorHandlers = map[string]MessageProcessorHandler{
	"response.text.delta": {
		Callback: handleTextDelta,
	},
	"response.content_part.done": {
		Callback: handleContentPartDone,
	},
	"response.function_call_arguments.done": {
		Callback: handleFunctionCallArgumentsDone,
	},
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
	requiredParams := map[string]reflect.Type{
		"arguments": reflect.TypeOf(""),
		"call_id":   reflect.TypeOf(""),
		"name":      reflect.TypeOf(""),
	}

	validated, err := validation.ValidateRequiredParams(message, requiredParams)
	if err != nil {
		return fmt.Errorf("function call validation failed: %w", err)
	}

	argumentsRaw := validated.GetString("arguments")
	var arguments map[string]any

	if err := json.Unmarshal([]byte(argumentsRaw), &arguments); err != nil {
		return fmt.Errorf("failed to unmarshal arguments: %w", err)
	}

	callId := validated.GetString("call_id")
	name := validated.GetString("name")

	if ok, returnEvent := functions.Handler(name, arguments, callId); ok {
		if err := p.client.SendMessage(returnEvent); err != nil {
			return fmt.Errorf("failed to send return event: %w", err)
		}

		responseCreate := builders.BuildResponseCreateMsg()
		if err := p.client.SendMessage(responseCreate); err != nil {
			return fmt.Errorf("failed to send response create message: %w", err)
		}
	}

	return nil
}
