package messages

import "fmt"

func handleTextDelta(messageProcessor *MessageProcessor, message map[string]interface{}) error {
	fmt.Print(message["delta"])

	return nil
}

func handleContentPartDone(messageProcessor *MessageProcessor, message map[string]interface{}) error {
	fmt.Println()

	return messageProcessor.client.StartUserGPTChat()
}

func handleFunctionCallArgumentsDone(messageProcessor *MessageProcessor, message map[string]interface{}) error {
	return messageProcessor.handleFunctionCall(message)
}
