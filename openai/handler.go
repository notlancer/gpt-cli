package openai

import (
	"encoding/json"
	"fmt"
	"github.com/notlancer/gpt-cli/openai/events"
	"github.com/notlancer/gpt-cli/openai/funcCall"
	"log"
)

type WsMessageHandler struct {
	Callback func(Client *Client, message map[string]interface{})
}

var wsMessagesHandlers = map[string]WsMessageHandler{
	"response.text.delta": {
		Callback: responseTextDeltaEvent,
	},
	"response.content_part.done": {
		Callback: contentPartDoneEvent,
	},
	"response.function_call_arguments.done": {
		Callback: responseFunCallArgEvent,
	},
}

func responseFunCallArgEvent(client *Client, message map[string]interface{}) {
	argumentsRaw := message["arguments"].(string)
	var arguments map[string]any

	err := json.Unmarshal([]byte(argumentsRaw), &arguments)
	if err != nil {
		log.Fatal("Error unmarshaling:", err)
		return
	}

	callId := message["call_id"].(string)

	if ok, returnEvent := funcCall.Handler(message["name"].(string), arguments, callId); ok {
		SendWsMessage(client.ws, returnEvent)

		responseCreate := events.BuildResponseCreateMsg()
		SendWsMessage(client.ws, responseCreate)
	}
}

func contentPartDoneEvent(client *Client, _ map[string]interface{}) {
	println()

	client.StartUserGPTChat()
}

func responseTextDeltaEvent(_ *Client, message map[string]interface{}) {
	fmt.Print(message["delta"])
}

func HandleWsMessage(client *Client, rawMessage []byte) {
	var message map[string]interface{}

	err := json.Unmarshal(rawMessage, &message)
	if err != nil {
		log.Fatal("Error unmarshaling:", err)
		return
	}

	if handler, ok := wsMessagesHandlers[message["type"].(string)]; ok {
		handler.Callback(client, message)
	}
}
