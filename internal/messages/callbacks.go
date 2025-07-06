package messages

import (
	"fmt"
	"reflect"

	"github.com/notlancer/gpt-cli/internal/validation"
)

func handleTextDelta(messageProcessor *MessageProcessor, message map[string]interface{}) error {
	requiredParams := map[string]reflect.Type{
		"delta": reflect.TypeOf(""),
	}

	validated, err := validation.ValidateRequiredParams(message, requiredParams)
	if err != nil {
		return fmt.Errorf("text delta validation failed: %w", err)
	}

	delta := validated.GetString("delta")
	fmt.Print(delta)

	return nil
}

func handleContentPartDone(messageProcessor *MessageProcessor, message map[string]interface{}) error {
	fmt.Println()

	return messageProcessor.client.StartUserGPTChat()
}

func handleFunctionCallArgumentsDone(messageProcessor *MessageProcessor, message map[string]interface{}) error {
	return messageProcessor.handleFunctionCall(message)
}
