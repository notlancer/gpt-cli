package openai

import (
	"encoding/json"
	"fmt"
)

type WsMessageHandler struct {
	Callback func(Client *Client, message map[string]interface{})
}

var wsMessagesHandlers = map[string]WsMessageHandler{
	"response.text.delta": {
		Callback: ResponseTextDelta,
	},
	"response.content_part.done": {
		Callback: ContentPartDone,
	},
}

func ContentPartDone(client *Client, _ map[string]interface{}) {
	println()
	RequestUserInput(client)
}

func ResponseTextDelta(_ *Client, message map[string]interface{}) {
	fmt.Print(message["delta"])
}

func HandleWsMessage(client *Client, rawMessage []byte) {
	var message map[string]interface{}

	err := json.Unmarshal(rawMessage, &message)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	if handler, ok := wsMessagesHandlers[message["type"].(string)]; ok {
		handler.Callback(client, message)
	}
}
