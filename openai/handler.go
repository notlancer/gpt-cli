package openai

import (
	"encoding/json"
	"fmt"
	"github.com/notlancer/gpt-cli/openai/events"
	"strconv"
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
	var arguments map[string]int

	err := json.Unmarshal([]byte(argumentsRaw), &arguments)
	if err != nil {
		fmt.Println("Error unmarshaling:", err)
		return
	}

	callId := message["call_id"].(string)
	var sum = arguments["number1"] * arguments["number2"]

	// i know it's could be float, wip!
	responseMsg := events.BuildConvCreateCallFuncMsg(callId, strconv.Itoa(sum))

	client.SendWsMessage(responseMsg)

	responseCreate := events.BuildResponseCreateMsg()
	client.SendWsMessage(responseCreate)

	fmt.Println("called func call")
}

func contentPartDoneEvent(client *Client, _ map[string]interface{}) {
	println()
	RequestUserInput(client)
}

func responseTextDeltaEvent(_ *Client, message map[string]interface{}) {
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
